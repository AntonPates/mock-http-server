package test

import (
	"context"
	"log"
	"net"
	"net/url"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	generalExitCode = 1
)

func TestMain(m *testing.M) {
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../build",
			Dockerfile: "Dockerfile",
		},
		Name:         "mock-http-server",
		ExposedPorts: []string{"8080/tcp"},
		WaitingFor: wait.ForAll(wait.ForListeningPort("8080/tcp"),
			wait.ForHTTP("/healthcheck").WithPort("8080/tcp").WithStatusCodeMatcher(func(status int) bool {
				return status == 200
			})),
	}
	ctx := context.Background()
	mockServer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		log.Println(err)
		os.Exit(generalExitCode)
	}

	mockServerHost, err := mockServer.Host(ctx)
	if err != nil {
		log.Println(err)
		os.Exit(generalExitCode)
	}

	mockServerPort, err := mockServer.MappedPort(ctx, "8080")
	if err != nil {
		log.Println(err)
		os.Exit(generalExitCode)
	}

	mockServerURL := (&url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(mockServerHost, mockServerPort.Port()),
	}).String()

	os.Setenv("MOCK_SERVER_URL", mockServerURL)

	exitCode := m.Run()
	mockServer.Terminate(ctx)
	os.Exit(exitCode)
}
