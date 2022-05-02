package server

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecognation(t *testing.T) {
	file, err := os.Open("/home/lingyin/Desktop/test.png")
	if err != nil {
		log.Fatal(err)
	}
	_, err = RecoFile(file)
	assert.Equal(t, nil, err, "failed.")
}

func TestStatus(t *testing.T) {
	Status()
}
