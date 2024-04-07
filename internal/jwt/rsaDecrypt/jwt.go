package rsaDecrypt

type Request struct {
	Jwt string `json:"jwt"`
}

type Response struct {
	Body interface{} `json:"jwt"`
}

type ResponseErr struct {
	Code        string `json:"code" example:"400"`
	Description string `json:"description" example:"bad request"`
}
