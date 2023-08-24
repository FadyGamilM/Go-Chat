package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	storedPass       = "fady123456789"
	correctLoginPass = "fady123456789"
	wrongLoginPass   = "fady123456"
)

func TestPasswordLogin(t *testing.T) {
	hashedPass, err := HashPassword(storedPass)
	require.NoError(t, err)

	match, err := CheckPassword(correctLoginPass, hashedPass)
	require.NoError(t, err)
	require.Equal(t, true, match)

	notMatch, err := CheckPassword(wrongLoginPass, hashedPass)
	require.Error(t, err)
	require.Equal(t, false, notMatch)
}
