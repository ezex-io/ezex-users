package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBcrypt_GenerateAndCompare(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "simple password",
			password: "hello123",
		},
		{
			name:     "strong password",
			password: "Tr0ub4dor&3!@#",
		},
		{
			name:     "empty password",
			password: "",
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			hashed, err := Generate(test.password)
			require.NoError(t, err, "Generate should not error")
			require.NotEmpty(t, hashed, "Hashed password should not be empty")

			match := Compare(test.password, hashed)
			require.True(t, match, "Compare should return true for correct password")
		})
	}
}

func TestBcrypt_MustHash_PanicsOnError(t *testing.T) {
	defer func() {
		r := recover()
		require.Nil(t, r, "MustHash should not panic for valid input")
	}()

	_ = MustHash("safePassword123")
}

func TestBcrypt_Compare_InvalidPassword(t *testing.T) {
	hash, err := Generate("original")
	require.NoError(t, err)

	match := Compare("wrongPassword", hash)
	require.False(t, match, "Compare should return false for incorrect password")
}
