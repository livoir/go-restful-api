package test

import (
	"belajar-golang-restful-api/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
