package crp

type ReqBody struct {
	Data string `json:"data"`
}

type RspBody struct {
	Data any `json:"data"`
}
