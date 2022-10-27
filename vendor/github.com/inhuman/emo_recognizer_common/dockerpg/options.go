package dockerpg

import (
	"crypto/rand"
	"math"
	"math/big"
)

const (
	defaultPgUser     = "postgres"
	defaultPgPassword = "123"
	defaultPgPort     = 5432
	defaultPbDB       = "postgres"
	defaultPgHost     = "localhost"
	defaultSslMode    = "disable"

	defaultMigrationPath = "file://./migrations"

	defaultContainerName   = "test_pg"
	defaultContainerRepo   = "core.harbor.k8s.devim.team/proxy/library/postgres"
	defaultContainerTag    = "11"
	defaultContainerExpire = 120 // seconds
)

type Opts struct {
	PgUser          string
	PgPassword      string
	PgPort          int
	PgHost          string
	PgDb            string //nolint:stylecheck // cannot be changed for backwards compatibility
	PgSslMode       string
	MigrationsPath  string
	ContainerName   string
	ContainerRepo   string
	ContainerTag    string
	ContainerExpire uint
	PortBind        bool // bind docker port to host port
}

func defaultOptions() *Opts {
	return &Opts{
		PgUser:          defaultPgUser,
		PgPassword:      defaultPgPassword,
		PgPort:          defaultPgPort,
		PgHost:          defaultPgHost,
		PgDb:            defaultPbDB,
		PgSslMode:       defaultSslMode,
		MigrationsPath:  defaultMigrationPath,
		ContainerName:   defaultContainerName,
		ContainerRepo:   defaultContainerRepo,
		ContainerTag:    defaultContainerTag,
		ContainerExpire: defaultContainerExpire,
	}
}

type Option func(o *Opts)

func WithPgUser(pgUser string) Option {
	return func(o *Opts) {
		o.PgUser = pgUser
	}
}

func WithPgPassword(pgPassword string) Option {
	return func(o *Opts) {
		o.PgPassword = pgPassword
	}
}

func WithPgPort(pgPort int) Option {
	return func(o *Opts) {
		o.PgPort = pgPort
	}
}

func WithPgHost(pgHost string) Option {
	return func(o *Opts) {
		o.PgHost = pgHost
	}
}

//nolint:stylecheck // correct func name
func WithPgDbName(pgDbName string) Option {
	return func(o *Opts) {
		o.PgDb = pgDbName
	}
}

func WithSslMode(pgSslMode string) Option {
	return func(o *Opts) {
		o.PgSslMode = pgSslMode
	}
}

// WithMigrationPath - add migration path with prefix,
// for example "file://./migrations"
func WithMigrationPath(migrationsPath string) Option {
	return func(o *Opts) {
		o.MigrationsPath = migrationsPath
	}
}

func WithContainerName(containerName string) Option {
	return func(o *Opts) {
		o.ContainerName = containerName
	}
}

func WithContainerRepo(containerRepo string) Option {
	return func(o *Opts) {
		o.ContainerRepo = containerRepo
	}
}

func WithContainerTag(containerTag string) Option {
	return func(o *Opts) {
		o.ContainerTag = containerTag
	}
}

func WithContainerExpire(expireSeconds uint) Option {
	return func(o *Opts) {
		o.ContainerExpire = expireSeconds
	}
}

const (
	minPgRandomPort = 10000
	maxPgRandomPort = 50000
)

func WithRandomPgPort() Option {
	return func(o *Opts) {
		n, _ := rand.Int(rand.Reader, big.NewInt(maxPgRandomPort-minPgRandomPort+1))
		port := float64(n.Int64() + minPgRandomPort)
		port = math.Sqrt(port * port)

		o.PgPort = int(port)
	}
}
