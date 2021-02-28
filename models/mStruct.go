package models

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) SetData(data interface{}) {
	r.Data = data
}

func (r *Response) SetMessage(message string) {
	r.Message = message
}

func (r *Response) SetError(error bool) {
	r.Error = error
}
