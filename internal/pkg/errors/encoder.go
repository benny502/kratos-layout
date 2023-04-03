package errors

import (
	"fmt"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type HTTPError struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Result  struct{} `json:"result"`
}

// Error implements error
func (err *HTTPError) Error() string {
	return fmt.Sprintf("Http Error: %d %s", err.Code, err.Message)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	return NewHTTPError(1, err.Error())
}

func ErrorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.Write(body)
}

func NewHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: msg,
		Result:  struct{}{},
	}
}
