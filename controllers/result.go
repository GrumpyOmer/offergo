package controllers

type result struct {
	code int
	msg  string
	data interface{}
}

//result:success
func (r *result) Success(data interface{}) result {
	result := result{
		code: 200,
		msg:  "ok",
		data: data,
	}
	return result
}

//result:error
func (r *result) Error(msg string) result {
	result := result{
		code: 400,
		msg:  msg,
		data: nil,
	}
	return result
}
