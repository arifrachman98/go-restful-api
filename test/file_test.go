package test

import (
	"testing"

	"github.com/arifrachman98/go-restful-api/simple"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	conn, clean := simple.InitConnection("Called Now")
	assert.NotNil(t, conn)

	clean()
}
