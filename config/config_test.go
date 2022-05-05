package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Builder(t *testing.T) {
	cfg := New("test_app", "test_service", "test_env").
		WithAudit("url_test", "user_test", "pass_test").
		WithEnvironment("pod_test", "node_test")

	assert.NotNil(t, cfg)
}
