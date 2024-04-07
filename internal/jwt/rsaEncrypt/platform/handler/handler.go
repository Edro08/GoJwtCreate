package handler

import (
	"GoJwtCreate/internal/jwt/rsaEncrypt"
	"GoJwtCreate/internal/jwt/rsaEncrypt/operations"
	"GoJwtCreate/kit/constants"
	"GoJwtCreate/kit/logger"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
)

const TitleTransportEncrypt = "---- ENCRYPT TRANSPORT ----"

type JwtHandler struct {
	jwtOperations operations.ICreate
	logger        logger.ILogger
}

func NewJwtEncryptHandler(jwtOperations operations.ICreate, logger logger.ILogger) JwtHandler {
	return JwtHandler{
		jwtOperations: jwtOperations,
		logger:        logger,
	}
}

func (h JwtHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	reqInterface, ctx, err := h.Decoder(r.Context(), r)

	if err != nil {
		h.EncoderError(ctx, w, err)
		return
	}

	req := reqInterface.(rsaEncrypt.Request)
	resp, err := h.jwtOperations.Create(ctx, req)

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

	var requestData rsaEncrypt.Request
	err := decoder.Decode(&requestData)
	if err != nil {
		return nil, ctxNew, err
	}

	return requestData, ctxNew, nil
}

func (h JwtHandler) Encoder(ctx context.Context, w http.ResponseWriter, response interface{}) {
	resp := response.(rsaEncrypt.Response)
	statusCode := http.StatusOK
	h.logger.Info(TitleTransportEncrypt, "Status Code", statusCode, "Response", resp)
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h JwtHandler) EncoderError(ctx context.Context, w http.ResponseWriter, response interface{}) {
	err := response.(error)
	var r rsaEncrypt.ResponseErr
	statusCode := http.StatusInternalServerError

	switch err {
	default:
		r = rsaEncrypt.ResponseErr{
			Code:        "500",
			Description: err.Error(),
		}
		statusCode = http.StatusInternalServerError
	}

	h.logger.Info(TitleTransportEncrypt, "Status Code", statusCode, "Response", r)

	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(r)
}
