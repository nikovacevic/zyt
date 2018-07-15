package err

import (
	"os"

	"github.com/nikovacevic/zyt/internal/pkg/log"
)

// Service is an ErrorService
type Service struct {
	logger *log.Logger
}

// NewService creates a new ErrorService
func NewService(l *log.Logger) *Service {
	return &Service{
		logger: l,
	}
}

// CheckAndFatal checks if an error is nil. If not, it logs the error and exits.
func (s *Service) CheckAndFatal(err error) {
	if err != nil {
		s.logger.Log(log.ERROR, err.Error())
		os.Exit(1)
	}
}

// CheckAndLog checks if an error is nil. If not, it logs the error.
func (s *Service) CheckAndLog(err error) {
	if err != nil {
		s.logger.Log(log.ERROR, err.Error())
	}
}

// CheckAndPanic checks if an error is nil. If not, it logs the error and panics.
func (s *Service) CheckAndPanic(err error) {
	if err != nil {
		s.logger.Log(log.ERROR, err.Error())
		panic(err)
	}
}
