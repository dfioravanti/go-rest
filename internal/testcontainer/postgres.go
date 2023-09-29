package testcontainer

import (
	"context"
	"fmt"

	"github.com/docker/go-connections/nat"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// PostgreSQLContainer wraps testcontainers.Container with extra methods.
type (
	PostgreSQLContainer struct {
		testcontainers.Container
		Config PostgreSQLContainerConfig
	}

	PostgreSQLContainerOption func(c *PostgreSQLContainerConfig)

	PostgreSQLContainerConfig struct {
		ImageTag   string
		User       string
		Password   string
		MappedPort string
		Database   string
		Host       string
	}
)

// GetDSN returns DB connection URL.
func (c PostgreSQLContainer) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Config.User, c.Config.Password, c.Config.Host, c.Config.MappedPort, c.Config.Database)
}

// NewPostgreSQLContainer creates and starts a PostgreSQL container.
func NewPostgreSQLContainer(ctx context.Context, opts ...PostgreSQLContainerOption) (*PostgreSQLContainer, error) {
	const (
		psqlImage = "postgres"
		psqlPort  = "5432"
	)

	// Define container ENVs
	config := PostgreSQLContainerConfig{
		ImageTag: "15",
		User:     "user",
		Password: "password",
		Database: "db_test",
	}
	for _, opt := range opts {
		opt(&config)
	}

	containerPort := psqlPort + "/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Env: map[string]string{
				"POSTGRES_USER":     config.User,
				"POSTGRES_PASSWORD": config.Password,
				"POSTGRES_DB":       config.Database,
			},
			ExposedPorts: []string{
				containerPort,
			},
			Image:      fmt.Sprintf("%s:%s", psqlImage, config.ImageTag),
			WaitingFor: wait.ForListeningPort("5432/tcp"),
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting request provider: %w", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting host for: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(containerPort))
	if err != nil {
		return nil, fmt.Errorf("getting mapped port for (%s): %w", containerPort, err)
	}
	config.MappedPort = mappedPort.Port()
	config.Host = host

	fmt.Println("Host:", config.Host, config.MappedPort)

	return &PostgreSQLContainer{
		Container: container,
		Config:    config,
	}, nil
}
