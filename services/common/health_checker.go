package common

type HealthChecker interface {
	Check() error
}
