package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)

	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Code, "not_found")
	assert.EqualValues(t, err.Message, "user 0 does not exist")
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "we were not expecting an error on user 123")
	assert.NotNil(t, user, "we were expecting a user with userId 123")

	assert.EqualValues(t, user.FirstName, "Leonardo")
	assert.EqualValues(t, user.LastName, "DiCaprio")
	assert.EqualValues(t, user.Email, "leonard@cap.io")
}
