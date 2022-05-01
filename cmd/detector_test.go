package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetector(t *testing.T) {
	assert.Equal(t, "", Detector(), "Detector() test failed.")
}
