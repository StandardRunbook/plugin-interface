package shared

import (
	"context"
	"errors"
	"testing"

	plugininterface "github.com/StandardRunbook/plugin-interface/hypothesis-interface/github.com/StandardRunbook/hypothesis"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockIPlugin is a mock implementation of the IPlugin interface
type MockIPlugin struct {
	mock.Mock
}

func (m *MockIPlugin) Init(config map[string]string) error {
	args := m.Called(config)
	return args.Error(0)
}

func (m *MockIPlugin) Name() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockIPlugin) Version() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockIPlugin) Run() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockIPlugin) ParseOutput() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestGRPCServer_Init(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	config := &plugininterface.Config{Parameters: map[string]string{"key": "value"}}
	mockPlugin.On("Init", config.GetParameters()).Return(nil)

	resp, err := server.Init(ctx, config)

	require.NoError(t, err)
	require.Equal(t, &plugininterface.InitResponse{ErrorMessage: ApplicationResponseSuccess}, resp)
}

func TestGRPCServer_Init_Error(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	config := &plugininterface.Config{Parameters: map[string]string{"key": "value"}}
	mockPlugin.On("Init", config.GetParameters()).Return(errors.New("init error"))

	resp, err := server.Init(ctx, config)

	require.Error(t, err)
	require.Nil(t, resp)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_Name(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("Name").Return("test-name", nil)

	resp, err := server.Name(ctx, &plugininterface.Empty{})

	require.NoError(t, err)
	require.Equal(t, "test-name", resp.Name)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_Name_Error(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("Name").Return("", errors.New("name error"))

	resp, err := server.Name(ctx, &plugininterface.Empty{})

	require.Error(t, err)
	require.Nil(t, resp)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_Version(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("Version").Return("1.0.0", nil)

	resp, err := server.Version(ctx, &plugininterface.Empty{})

	require.NoError(t, err)
	require.Equal(t, "1.0.0", resp.Version)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_Version_Error(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("Version").Return("", errors.New("version error"))

	resp, err := server.Version(ctx, &plugininterface.Empty{})

	require.Error(t, err)
	require.Nil(t, resp)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_Run(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("Run").Return(nil)

	resp, err := server.Run(ctx, &plugininterface.Empty{})

	require.NoError(t, err)
	require.Equal(t, &plugininterface.RunResponse{ErrorMessage: ApplicationResponseSuccess}, resp)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_Run_Error(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("Run").Return(errors.New("run error"))

	resp, err := server.Run(ctx, &plugininterface.Empty{})

	require.Error(t, err)
	require.Nil(t, resp)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_ParseOutput(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("ParseOutput").Return("test-output", nil)

	resp, err := server.ParseOutput(ctx, &plugininterface.Empty{})

	require.NoError(t, err)
	require.Equal(t, "test-output", resp.Output)
	mockPlugin.AssertExpectations(t)
}

func TestGRPCServer_ParseOutput_Error(t *testing.T) {
	mockPlugin := new(MockIPlugin)
	server := &GRPCServer{Impl: mockPlugin}
	ctx := context.Background()

	mockPlugin.On("ParseOutput").Return("", errors.New("parse error"))

	resp, err := server.ParseOutput(ctx, &plugininterface.Empty{})

	require.Error(t, err)
	require.Nil(t, resp)
	mockPlugin.AssertExpectations(t)
}
