package response

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
	Error      interface{}
}

func MakeResponse(stauscode int, message string, data, error interface{}) *Response {
	return &Response{
		StatusCode: stauscode,
		Message:    message,
		Data:       data,
		Error:      error,
	}
}

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	PhNo    string `json:"ph_no"`
	Gender  string `json:"gender"`
	City    string `json:"city"`
	Country string `json:"country"`
	IsBlock bool   `json:"is_block"`
}
