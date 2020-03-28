package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelp(t *testing.T) {
	assert.NoError(t, Help())
}
