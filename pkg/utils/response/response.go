package response



type Response struct{
	StatusCode int
	Message string
	Data interface{}
	Error interface{}
}

func MakeResponse(stauscode int,message string,data,error interface{})*Response{
	return &Response{
		StatusCode: stauscode,
		Message: message,
		Data: data,
		Error: error,
	}
}