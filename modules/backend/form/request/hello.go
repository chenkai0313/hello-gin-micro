package form

type SayHelloReq struct {
	Name    string `json:"name"  validate:"required"`
	Content string `json:"content" validate:"required"`
}
