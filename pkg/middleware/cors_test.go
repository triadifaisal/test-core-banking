package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware(t *testing.T) {
	got := CORSMiddleware()
	assert.NotNil(t, got)
}
