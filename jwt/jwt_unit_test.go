// +build unit

package jwt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x64puzzle/go-common/jwt"
)

func TestJWT(t *testing.T) {
	token, err := jwt.Generate(map[string]string{
		"username": "semir",
		"email":    "semir@mail.com",
		"id":       "semir-123",
	})

	assert.NoError(t, err, "Err occured: ", err)
	assert.NotEmpty(t, token, "Token should not be empty")

	claims, valid := jwt.ValidateAndExtract(token)
	assert.True(t, valid, "Token should be valid")
	assert.NotEmpty(t, claims, "Claims empty")
}
