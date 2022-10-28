package pgx

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strings"
	"time"
)

type Pool interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(ctx context.Context, sql string, args []interface{},
		scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
	Config() *pgxpool.Config
}

type PgArgs struct {
	Host      string
	Port      int
	DB        string
	User      string
	Password  string
	SslMode   string
	Additions map[string]string
}

func (a *PgArgs) String() (string, error) {
	if a.Host == "" {
		return "", fmt.Errorf("host is empty")
	}

	if a.Port == 0 {
		return "", fmt.Errorf("port is empty")
	}

	if a.DB == "" {
		return "", fmt.Errorf("database is empty")
	}

	if a.User == "" {
		return "", fmt.Errorf("user is empty")
	}

	if a.Password == "" {
		return "", fmt.Errorf("password is empty")
	}

	ssl := a.SslMode
	if a.SslMode == "" {
		ssl = "disable"
	}

	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		a.Host, a.Port, a.DB, a.User, a.Password, ssl)

	if a.Additions != nil {
		for key, value := range a.Additions {
			connString += fmt.Sprintf(" %s=%s", key, value)
		}
	}

	return connString, nil
}

func QueryFromFile(pgxPool Pool, filePath string) error {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error read file for query: %w", err)
	}

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = pgxPool.Exec(ctxWithTimeout, string(fileBytes))
	if err != nil {
		return fmt.Errorf("error exec query from file: %w", err)
	}

	return nil
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

func ToInClause(t ...fmt.Stringer) string {
	inClause := ""

	for i := range t {
		inClause += fmt.Sprintf(`"%s",`, t[i])
	}

	return strings.TrimRight(inClause, ",")
}
