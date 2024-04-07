package operations

import (
	"GoJwtCreate/internal/jwt/rsaEncrypt"
	"context"
)

type ICreate interface {
	Create(ctx context.Context, request rsaEncrypt.Request) (rsaEncrypt.Response, error)
}
