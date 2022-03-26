package password

import (
	"github.com/0RAJA/Road/internal/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashPassword(t *testing.T) {
	//testHashPassword()
	password := utils.RandomString(10)

	hashPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotZero(t, hashPassword)
	require.NoError(t, CheckPassword(password, hashPassword))
	wrongPassword := utils.RandomString(10)
	require.Error(t, CheckPassword(wrongPassword, hashPassword))
}
