package logger_test

import (
	"fmt"
	"testing"

	"github.com/eduardohslfreire/animalia-api/config/env"
	"github.com/eduardohslfreire/animalia-api/pkg/logger"
	"github.com/stretchr/testify/assert"
)

const (
	testDebug = "[TEST] - Debug: %s"
	testInfo  = "[TEST] - Info: %s"
	testWarn  = "[TEST] - Warn: %s"
	testError = "[TEST] - Error: %s"
)

func TestLogger(t *testing.T) {
	t.Run("success-severity", func(t *testing.T) {
		env.LogLevel = "DEBUG"

		logger := logger.NewGenericLogger()
		logger.SetContext("teste")

		logger.LogIt("DEBUG", fmt.Sprintf(testDebug, "teste"), nil)
		logger.LogIt("INFO", fmt.Sprintf(testInfo, "teste"), nil)
		logger.LogIt("WARN", fmt.Sprintf(testWarn, "teste"), nil)
		logger.LogIt("ERROR", fmt.Sprintf(testError, "teste"), nil)

		assert.NotNil(t, logger)
	})
	t.Run("success-loglevel-debug", func(t *testing.T) {
		env.LogLevel = "DEBUG"

		logger := logger.NewGenericLogger()
		logger.SetContext("teste")

		logger.LogIt("DEBUG", fmt.Sprintf(testDebug, "teste"), nil)

		assert.NotNil(t, logger)
	})
	t.Run("success-loglevel-info", func(t *testing.T) {
		env.LogLevel = "INFO"

		logger := logger.NewGenericLogger()
		logger.SetContext("teste")

		logger.LogIt("INFO", fmt.Sprintf(testInfo, "teste"), nil)

		assert.NotNil(t, logger)
	})
	t.Run("success-loglevel-warn", func(t *testing.T) {
		env.LogLevel = "WARN"

		logger := logger.NewGenericLogger()
		logger.SetContext("teste")

		logger.LogIt("WARN", fmt.Sprintf(testWarn, "teste"), nil)

		assert.NotNil(t, logger)
	})

	t.Run("success-loglevel-disable", func(t *testing.T) {
		logger := logger.NewLogger()
		logger.SetContext("teste")

		logger.LogIt("", fmt.Sprintf(testDebug, "teste"), nil)

		assert.NotNil(t, logger)
	})

	t.Run("success-loglevel-default", func(t *testing.T) {
		env.LogLevel = ""
		logger := logger.NewLogger()
		logger.SetContext("teste")

		logger.LogIt("", fmt.Sprintf(testDebug, "teste"), nil)

		assert.NotNil(t, logger)
	})

	t.Run("success-field", func(t *testing.T) {
		field := map[string]interface{}{"teste": ""}

		logger := logger.NewGenericLogger()
		logger.SetContext("teste")

		logger.LogIt("WARN", fmt.Sprintf(testWarn, "teste"), field)

		assert.NotNil(t, logger)
	})

	t.Run("error-root", func(t *testing.T) {
		env.LogLevel = "ERROR"

		logger := logger.NewGenericLogger()
		logger.SetContext("teste")

		logger.LogIt("ERROR", fmt.Sprintf(testError, "teste"), nil)
		logger = nil
		assert.Nil(t, logger)
	})
}
