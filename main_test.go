package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Nil(t, nil)
}

func TestLoadConfig(t *testing.T) {
	t.Setenv("REDIRECT_TARGET", "")
}
