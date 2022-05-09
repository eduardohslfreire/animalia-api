package env_test

import (
	"os"
	"testing"

	"github.com/eduardohslfreire/animalia-api/config/env"
	"github.com/stretchr/testify/assert"
)

func TestEnvConfig(t *testing.T) {
	os.Setenv("Host", "Host")
	assert.NotNil(t, env.DbHost)
}
