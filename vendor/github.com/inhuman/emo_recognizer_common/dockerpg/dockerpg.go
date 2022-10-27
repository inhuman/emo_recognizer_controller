//nolint:gomnd,wrapcheck,gci // false positive warnings
package dockerpg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"

	// drivers needed
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type PgLogger struct {
	z *zap.Logger
}

func NewPgLogger(z *zap.Logger) *PgLogger {
	return &PgLogger{z: z}
}

func (l *PgLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	var fields []zap.Field

	for name, value := range data {
		fields = append(fields, zap.Any(name, value))
	}

	l.z.Info(msg, fields...)
}

type DockerPg struct {
	pool     *dockertest.Pool
	resource *dockertest.Resource
	options  *Opts
	network  *dockertest.Network
}

func (d *DockerPg) Close() error {
	return d.pool.Purge(d.resource)
}

func (d *DockerPg) AppliedOptions() Opts {
	return *d.options
}

func (d *DockerPg) Network() *dockertest.Network {
	return d.network
}

func (d *DockerPg) Pool() *dockertest.Pool {
	return d.pool
}

func (d *DockerPg) Resource() *dockertest.Resource {
	return d.resource
}

func NewDockerPg(opts ...Option) (*DockerPg, error) {
	dockerPg := &DockerPg{}

	options := defaultOptions()

	for _, opt := range opts {
		opt(options)
	}

	var err error

	dockerPg.pool, err = dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not connect to docker: %w", err)
	}

	res, ok := dockerPg.pool.ContainerByName(options.ContainerName)
	if ok {
		err = dockerPg.pool.Purge(res)
		if err != nil {
			return nil, fmt.Errorf("cannot purge previous pg instance: %w", err)
		}
	}

	results, err := dockerPg.pool.Client.PruneNetworks(docker.PruneNetworksOptions{
		Filters: map[string][]string{
			"label": []string{"name=docker_pg_network"},
		},
		Context: context.Background(),
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("remove network result: ", results)

	network, err := dockerPg.pool.CreateNetwork("docker_pg_network", func(config *docker.CreateNetworkOptions) {
		config.Labels = map[string]string{"name": "docker_pg_network"}
	})
	if err != nil {
		return nil, fmt.Errorf("can not create docker pg network")
	}

	runOpts := prepareRunOptions(options, network)

	dockerPg.network = network

	// pulls an image, creates a Container based on it and runs it
	dockerPg.resource, err = dockerPg.pool.RunWithOptions(runOpts, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped Container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		return nil, fmt.Errorf("could not start resource: %w", err)
	}

	// nolint:nosprintfhostport // тут делается дсн для базы а не просто хост порт
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		options.PgUser, options.PgPassword, options.PgHost, options.PgPort, options.PgDb, options.PgSslMode)

	log.Println("Connecting to database on url: ", databaseURL)

	err = dockerPg.resource.Expire(options.ContainerExpire) // Tell docker to hard kill the Container in 120 seconds
	if err != nil {
		return nil, fmt.Errorf("error run expire func: %w", err)
	}

	// exponential backoff-retry, because the application in the Container might not be ready to accept connections yet
	dockerPg.pool.MaxWait = 120 * time.Second
	if err = dockerPg.pool.Retry(func() error {
		var db *sql.DB
		if db, err = sql.Open("postgres", databaseURL); err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %w", err)
	}

	log.Println("run migrations")

	migrator, err := migrate.New(options.MigrationsPath, databaseURL)
	if err != nil {
		return nil, err
	}

	err = migrator.Up()
	if err != nil {
		return nil, err
	}

	dockerPg.options = options

	return dockerPg, nil
}

func prepareRunOptions(dockerPgOpts *Opts, network *dockertest.Network) *dockertest.RunOptions {
	runOptions := &dockertest.RunOptions{
		Networks:   []*dockertest.Network{network},
		Name:       dockerPgOpts.ContainerName,
		Repository: dockerPgOpts.ContainerRepo,
		Tag:        dockerPgOpts.ContainerTag,
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", dockerPgOpts.PgPassword),
			fmt.Sprintf("POSTGRES_USER=%s", dockerPgOpts.PgUser),
			fmt.Sprintf("POSTGRES_DB=%s", dockerPgOpts.PgDb),
			"listen_addresses='*'",
		},
	}

	appPort := docker.Port("5432/tcp")
	runOptions.PortBindings = map[docker.Port][]docker.PortBinding{
		appPort: {
			{
				HostIP:   "0.0.0.0",
				HostPort: fmt.Sprintf("%d", dockerPgOpts.PgPort),
			},
		},
	}

	return runOptions
}
