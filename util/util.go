package util

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

// Regexphone 判断手机号码是否正确
func Regexphone(phone string) bool {
	right, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, phone)
	return right
}

