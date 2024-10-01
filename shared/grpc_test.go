package shared

import (
	"context"
	"github.com/stretchr/testify/require"
	"net"
	"testing"

	plugininterface "github.com/StandardRunbook/plugin-interface/hypothesis-interface/github.com/StandardRunbook/hypothesis"
	"google.golang.org/grpc"
)

type MockServer struct {
	plugininterface.UnimplementedHypothesisServer
}

func (m *MockServer) Init(ctx context.Context, config *plugininterface.Config) (*plugininterface.InitResponse, error) {
	return &plugininterface.InitResponse{
		ErrorMessage: ApplicationResponseSuccess,
	}, nil
}

func (m *MockServer) Name(ctx context.Context, empty *plugininterface.Empty) (*plugininterface.NameResponse, error) {
	return &plugininterface.NameResponse{
		Name: "mock_server",
	}, nil
}

func (m *MockServer) Version(ctx context.Context, empty *plugininterface.Empty) (*plugininterface.VersionResponse, error) {
	return &plugininterface.VersionResponse{
		Version: "v1.0.0",
	}, nil
}

func (m *MockServer) Run(ctx context.Context, empty *plugininterface.Empty) (*plugininterface.RunResponse, error) {
	return &plugininterface.RunResponse{
		ErrorMessage: ApplicationResponseSuccess,
	}, nil
}

func (m *MockServer) ParseOutput(ctx context.Context, empty *plugininterface.Empty) (*plugininterface.ParseOutputResponse, error) {
	return &plugininterface.ParseOutputResponse{
		Output: "success",
	}, nil
}

func makeGRPCServer(t *testing.T) (*grpc.Server, string) {
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	// Register the mock server
	plugininterface.RegisterHypothesisServer(server, &MockServer{})

	go func() {
		err := server.Serve(listener)
		require.NoError(t, err)
	}()

	return server, listener.Addr().String()
}

func TestGRPCClient_HappyPath(t *testing.T) {
	server, addr := makeGRPCServer(t)
	defer server.Stop()

	// Dial the test server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	client := &GRPCClient{
		client: plugininterface.NewHypothesisClient(conn),
	}

	// Call the Init method
	err = client.Init(map[string]string{"key": "value"})
	require.NoError(t, err)

	// Call the Name method
	name, err := client.Name()
	require.NoError(t, err)
	require.Equal(t, "mock_server", name)

	// Call the Version method
	version, err := client.Version()
	require.NoError(t, err)
	require.Equal(t, "v1.0.0", version)

	// Call the Run method
	err = client.Run()
	require.NoError(t, err)

	// Call the ParseOutput method
	output, err := client.ParseOutput()
	require.NoError(t, err)
	require.Equal(t, "success", output)
}
