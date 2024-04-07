package operations

import (
	"GoJwtCreate/internal/jwt/rsaDecrypt"
	"GoJwtCreate/kit/config"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCreated struct {
	config config.IConfig
}

func NewJwtCreated(config config.IConfig) JwtCreated {
	return JwtCreated{
		config: config,
	}
}

func (j JwtCreated) Decrypt(ctx context.Context, request rsaDecrypt.Request) (rsaDecrypt.Response, error) {

	// Clave pública RSA como string
	publicKeyStr, _ := j.config.GetString("service.config.jwt.secrets.public")

	// Parsea la clave pública RSA
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
	if err != nil {
		fmt.Println("Error al parsear la clave pública:", err)
		return rsaDecrypt.Response{}, nil
	}

	// Verifica el token con la clave pública RSA
	parsedToken, err := jwt.Parse(request.Jwt, func(token *jwt.Token) (interface{}, error) {
		//Verificar firma de algoritmo
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("invalid signature algorithm")
		}
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("Error al verificar el token:", err)
		return rsaDecrypt.Response{}, nil
	}

	if parsedToken.Valid {
		fmt.Println("El token JWT es válido.")
	} else {
		fmt.Println("El token JWT no es válido.")
	}

	// Obtiene los claims del token
	claims, _ := parsedToken.Claims.(jwt.MapClaims)

	// Construye un mapa de tipo interface{} a partir de los claims
	body := make(map[string]interface{})
	for key, value := range claims {
		body[key] = value
	}

	return rsaDecrypt.Response{Body: body}, nil
}
