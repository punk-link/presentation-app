package common

import (
	"context"
	"sync"
	"time"

	"github.com/punk-link/logger"
	contracts "github.com/punk-link/presentation-contracts"
	"github.com/samber/do"
)

type HealthCheckService struct {
	grpcClient contracts.PresentationClient
	logger     logger.Logger
	mutex      sync.Mutex
}

func NewHealthCheckService(injector *do.Injector) (*HealthCheckService, error) {
	grpcClient := do.MustInvoke[contracts.PresentationClient](injector)
	logger := do.MustInvoke[logger.Logger](injector)

	service := &HealthCheckService{
		grpcClient: grpcClient,
		logger:     logger,
	}

	go service.checkConnection()

	return service, nil
}

func (t *HealthCheckService) Check() error {
	if !_isConnectionStable {
		return _connectionStatusError
	}

	return nil
}

func (t *HealthCheckService) checkConnection() {
	ticker := time.NewTicker(GRPC_CONNECTION_CHECK_TIMEOUT)
	for range ticker.C {
		_, err := t.grpcClient.CheckHealth(context.Background(), &contracts.HealthCheckRequest{})
		t.setConnectionStatus(err)
	}
}

func (t *HealthCheckService) setConnectionStatus(err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if err == nil {
		_isConnectionStable = true
		return
	}

	_isConnectionStable = false
	_connectionStatusError = err
	t.logger.LogWarn("gRPC connection degraded: %v", err)
}

var _isConnectionStable bool
var _connectionStatusError error

const GRPC_CONNECTION_CHECK_TIMEOUT = 5 * time.Second
