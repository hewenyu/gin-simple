package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

// 用户uuid
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
