package macro

import (
	"encoding/json"
	"net/http"
	"protoserver-go/pkg/model"
	"protoserver-go/pkg/proto/sys"
	"time"
)

func MacrosForRequestBody(macros []string, body []byte) []byte {
	return body
}

func MacrosForResponseBody(macros []string, body []byte) []byte {
	return body
}

func MacroStdResponse(macros []string, body []byte, r *http.Request) []byte {
	if ArrayContains(macros, sys.MACRO_STD_RESPONSE) {
		now := time.Now()
		apiresp := model.ApiResponse{
			Payload:   body,
			Timestamp: now.String(),
			TraceUUID: r.Header.Get("traceUUID"),
		}
		apiRespBytes, err := json.Marshal(apiresp)
		if err != nil {
			apiRespBytes = apiresp.StdErr(err)
		}
		// log.Println("........ apiRespBytes: ", string(apiRespBytes))
		return apiRespBytes
		// str, _ := strconv.Unquote(string(apiRespBytes))
		// return []byte(str)
	}
	return body
}

func ArrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
