package operations

import (
	"GoJwtCreate/internal/jwt/rsaEncrypt"
	"GoJwtCreate/kit/config"
	"context"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtCreated struct {
	config config.IConfig
}

func NewJwtCreated(config config.IConfig) JwtCreated {
	return JwtCreated{
		config: config,
	}
}

func (j JwtCreated) Create(ctx context.Context, request rsaEncrypt.Request) (rsaEncrypt.Response, error) {

	// Clave privada RSA como string
	privateKeyStr, _ := j.config.GetString("service.config.jwt.secrets.private")

	// Parsea la clave privada RSA
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyStr))
	if err != nil {
		fmt.Println("Error al parsear la clave privada:", err)
		return rsaEncrypt.Response{}, nil
	}

	// Crea un nuevo token JWT
	token := jwt.New(jwt.SigningMethodRS256)

	serverName, _ := j.config.GetString("server.name")

	// Define los claims del token (pueden ser personalizados según tus necesidades)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = serverName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Tiempo de expiración en 24 horas

	for key, value := range request.Payload {
		claims[key] = value
	}

	// Genera el token firmado con la clave secreta
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("Error al firmar el token:", err)
		return rsaEncrypt.Response{}, nil
	}

	fmt.Println("Token JWT:", tokenString)

	return rsaEncrypt.Response{
		Jwt: tokenString,
	}, nil
}
