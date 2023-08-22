package _type

type ErrorResponseDTO struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponseDTO) Error() string {
	return e.Message
}

// NewSomething create new instance of Something
func (e *ErrorResponseDTO) New(code string, message ...string) *ErrorResponseDTO {
	e.Status = "error"
	e.Code = code
	if len(message) > 0 {
		e.Message = message[0]
	}
	return e
}

type ResponseDto struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponseInf interface {
	Success(data interface{}, message ...string) *ResponseDto
	Error(message ...string) *ResponseDto
	Warning(message ...string) *ResponseDto
}

type response struct{}

func Response() ResponseInf {
	return &response{}
}

func (e *ResponseDto) Error() string {
	return e.Message
}

func (t *response) Success(data interface{}, message ...string) *ResponseDto {
	var msg = ""
	if len(message) > 0 {
		msg = message[0]
	}
	return &ResponseDto{
		Message: msg,
		Data:    data,
	}
}

func (t *response) Error(message ...string) *ResponseDto {
	var msg = ""
	if len(message) > 0 {
		msg = message[0]
	}
	return &ResponseDto{
		Message: msg,
	}
}

func (t *response) Warning(message ...string) *ResponseDto {
	var msg = ""
	if len(message) > 0 {
		msg = message[0]
	}
	return &ResponseDto{
		Message: msg,
	}
}
