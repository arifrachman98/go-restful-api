package test

import (
	"fmt"
	"testing"

	"github.com/arifrachman98/go-restful-api/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	fmt.Println(err)
	fmt.Println(simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpServ, err := simple.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpServ)
}

func TestSimpleServiceError(t *testing.T) {
	simpService, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpService)
}
