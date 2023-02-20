package request

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

// CustomClaims Custom claims structure
type CustomClaims struct {
	UUID       uuid.UUID `json:"UUID"`
	ID         uint      `json:"ID"`
	Username   string    `json:"username"`
	BufferTime int64     `json:"bufferTime"`
	jwt.RegisteredClaims
}

//type CustomClaimsRole struct {
//	*CustomClaims
//	Roles []string `json:"roles,omitempty"`
//}
