package rsaEncrypt

type Request struct {
	Payload map[string]interface{} `json:"jwt"`
}

type Response struct {
	Jwt string `json:"jwt"`
}

type ResponseErr struct {
	Code        string `json:"code" example:"400"`
	Description string `json:"description" example:"bad request"`
}
