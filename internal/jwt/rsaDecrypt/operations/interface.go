package operations

import (
	"GoJwtCreate/internal/jwt/rsaDecrypt"
	"context"
)

type IDecrypt interface {
	Decrypt(ctx context.Context, request rsaDecrypt.Request) (rsaDecrypt.Response, error)
}
