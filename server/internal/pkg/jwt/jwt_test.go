package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	"server/internal/model/sysManagement"
)

func init() {
	// Ensure signing key is set for tests
	global.TD27_CONFIG.JWT.SigningKey = "test-signing-key-for-unit-tests"
}

func TestNewJWT(t *testing.T) {
	j := NewJWT()
	require.NotNil(t, j)
	require.Equal(t, []byte(global.TD27_CONFIG.JWT.SigningKey), j.SigningKey)
}

func TestCreateTokenAndParseToken(t *testing.T) {
	j := NewJWT()
	require.NotNil(t, j)

	claims := sysManagement.CustomClaims{
		ID:         1,
		Username:   "admin",
		BufferTime: 3600,
		Roles: []sysManagement.RoleInfo{
			{ID: 1, RoleName: "admin"},
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t.Run("create and parse valid token", func(t *testing.T) {
		tokenString, err := j.CreateToken(claims)
		require.NoError(t, err)
		require.NotEmpty(t, tokenString)

		parsed, err := j.ParseToken(tokenString)
		require.NoError(t, err)
		require.NotNil(t, parsed)
		require.Equal(t, claims.ID, parsed.ID)
		require.Equal(t, claims.Username, parsed.Username)
		require.Len(t, parsed.Roles, 1)
		require.Equal(t, claims.Roles[0].ID, parsed.Roles[0].ID)
		require.Equal(t, claims.Roles[0].RoleName, parsed.Roles[0].RoleName)
	})

	t.Run("parse malformed token", func(t *testing.T) {
		_, err := j.ParseToken("not.a.token")
		require.ErrorIs(t, err, TokenMalformed)
	})

	t.Run("parse invalid token", func(t *testing.T) {
		_, err := j.ParseToken("invalid-token-string")
		require.Error(t, err)
	})

	t.Run("parse expired token", func(t *testing.T) {
		expiredClaims := sysManagement.CustomClaims{
			ID:       2,
			Username: "expired",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
			},
		}
		tokenString, err := j.CreateToken(expiredClaims)
		require.NoError(t, err)

		_, err = j.ParseToken(tokenString)
		require.ErrorIs(t, err, TokenExpired)
	})

	t.Run("parse token with wrong signing key", func(t *testing.T) {
		wrongJWT := &JWT{SigningKey: []byte("wrong-key")}
		tokenString, err := wrongJWT.CreateToken(claims)
		require.NoError(t, err)

		_, err = j.ParseToken(tokenString)
		require.Error(t, err)
	})
}

func TestCreateTokenByOldToken(t *testing.T) {
	j := NewJWT()
	require.NotNil(t, j)

	claims := sysManagement.CustomClaims{
		ID:       3,
		Username: "test",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	t.Run("create by old token", func(t *testing.T) {
		oldToken := "old-token-123"
		tokenString, err := j.CreateTokenByOldToken(oldToken, claims)
		require.NoError(t, err)
		require.NotEmpty(t, tokenString)

		parsed, err := j.ParseToken(tokenString)
		require.NoError(t, err)
		require.Equal(t, claims.ID, parsed.ID)
		require.Equal(t, claims.Username, parsed.Username)
	})

	t.Run("same old token returns same result within concurrency control", func(t *testing.T) {
		oldToken := "old-token-456"
		token1, err := j.CreateTokenByOldToken(oldToken, claims)
		require.NoError(t, err)

		token2, err := j.CreateTokenByOldToken(oldToken, claims)
		require.NoError(t, err)

		require.Equal(t, token1, token2)
	})
}
