package env_test

import (
	"testing"

	"github.com/eduardohslfreire/animalia-api/pkg/env"
	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	env := env.NewEnv()
	env.GetString("test")
	env.GetInt("test")
	env.GetBool("test")
	assert.NotNil(t, env)
}
