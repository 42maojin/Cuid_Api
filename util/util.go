package util

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// ResponseJSON 返回json格式和接收errorcode
func ResponseJSON(w http.ResponseWriter, errorCode int, msg string, v interface{}) error {
	r := map[string]interface{}{
		"msg":        msg,
		"error_code": errorCode,
		"data":       v,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return nil
}
