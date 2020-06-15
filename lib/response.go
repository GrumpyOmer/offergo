package lib

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//success 200
func (r *Response) Success(data interface{}) (response Response) {
	response = Response{
		200,
		"ok",
		data,
	}
	return
}

//error 400
func (r *Response) Error(msg string) (response Response) {
	response = Response{
		400,
		msg,
		nil,
	}
	return
}
