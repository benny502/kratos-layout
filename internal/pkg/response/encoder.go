package response

import (
	"encoding/json"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func ResponseEncoder(writer nethttp.ResponseWriter, request *nethttp.Request, response interface{}) error {

	if response == nil {
		return nil
	}

	codec, _ := http.CodecForRequest(request, "Accept")

	body, err := codec.Marshal(response)
	if err != nil {
		return err
	}
	var reply interface{}
	json.Unmarshal(body, &reply)

	se := FromResposne(reply)
	body, err = json.Marshal(se)

	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "application/"+codec.Name())
	_, err = writer.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func FromResposne(response interface{}) *HttpResponse {

	return &HttpResponse{
		Code:    0,
		Message: "",
		Result:  response,
	}
}
