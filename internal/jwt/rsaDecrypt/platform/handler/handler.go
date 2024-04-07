package handler

import (
	"GoJwtCreate/internal/jwt/rsaDecrypt"
	"GoJwtCreate/internal/jwt/rsaDecrypt/operations"
	"GoJwtCreate/kit/constants"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type JwtHandler struct {
	jwtOperations operations.IDecrypt
}

func NewJwtDecryptHandler(jwtOperations operations.IDecrypt) JwtHandler {
	return JwtHandler{
		jwtOperations: jwtOperations,
	}
}

func (h JwtHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	reqInterface, ctx, err := h.Decoder(r.Context(), r)

	if err != nil {
		h.EncoderError(ctx, w, err)
		return
	}

	req := reqInterface.(rsaDecrypt.Request)
	resp, err := h.jwtOperations.Decrypt(ctx, req)

	if err != nil {
		h.EncoderError(ctx, w, err)
		return
	}

	h.Encoder(ctx, w, resp)
}

func (h JwtHandler) Decoder(ctx context.Context, r *http.Request) (interface{}, context.Context, error) {
	processID := uuid.New()
	ctxNew := context.WithValue(ctx, constants.UUID, processID.String())

	ip := r.RemoteAddr
	ctxNew = context.WithValue(ctxNew, constants.IP, ip)

	decoder := json.NewDecoder(r.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	var requestData rsaDecrypt.Request
	err := decoder.Decode(&requestData)
	if err != nil {
		return nil, ctxNew, err
	}

	return requestData, ctxNew, nil
}

func (h JwtHandler) Encoder(ctx context.Context, w http.ResponseWriter, response interface{}) {
	resp := response.(rsaDecrypt.Response)
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h JwtHandler) EncoderError(ctx context.Context, w http.ResponseWriter, response interface{}) {
	err := response.(error)
	var r rsaDecrypt.ResponseErr

	switch err {
	default:
		r = rsaDecrypt.ResponseErr{
			Code:        "500",
			Description: err.Error(),
		}
		w.WriteHeader(500)
	}

	_ = json.NewEncoder(w).Encode(r)
}
