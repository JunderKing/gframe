package http

import (
	"fmt"
	"gframe/pkg/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CodeOk        = 0   // 成功
	CodeErr       = 100 // 错误
	CodeException = 500 // 异常
	CodeNoAuth    = 101 // 无权限
	CodeNotLogin  = 102 // 未登录
)

type RespStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespOk(data interface{}) *RespStruct {
	resp := &RespStruct{Code: CodeOk, Msg: "ok", Data: data}
	panic(resp)
}

func RespErr(msg string) {
	resp := &RespStruct{Code: CodeErr, Msg: msg}
	panic(resp)
}

func RespCode(code int, msg string, data interface{}) {
	resp := &RespStruct{Code: code, Msg: msg, Data: data}
	panic(resp)
}

func Validate(c *gin.Context, params interface{}) {
	if err := c.ShouldBind(params); err != nil {
		RespErr(err.Error())
	}
}

func CurlGet(urlString string, urlParams url.Values, msTimeout int, header http.Header) (*http.Response, []byte, error) {
	startTime := time.Now().UnixNano()
	client := http.Client{
		Timeout: time.Duration(msTimeout) * time.Millisecond,
	}
	urlString = AddGetDataToUrl(urlString, urlParams)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		logger.Warn("CurlFailed", map[string]interface{}{
			"url":       urlString,
			"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
			"method":    "GET",
			"args":      urlParams,
			"err":       err.Error(),
		})
		return nil, nil, err
	}
	if len(header) > 0 {
		req.Header = header
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Warn("CurlFailed", map[string]interface{}{
			"url":       urlString,
			"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
			"method":    "GET",
			"args":      urlParams,
			"err":       err.Error(),
		})
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("CurlFailed", map[string]interface{}{
			"url":       urlString,
			"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
			"method":    "GET",
			"args":      urlParams,
			"result":    Substr(string(body), 0, 1024),
			"err":       err.Error(),
		})
		return nil, nil, err
	}
	logger.Info("CurlSuccess", map[string]interface{}{
		"url":       urlString,
		"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		"method":    "GET",
		"args":      urlParams,
		"result":    Substr(string(body), 0, 1024),
	})
	return resp, body, nil
}

func CurlPost(urlString string, urlParams url.Values, msTimeout int, header http.Header, contextType string) (*http.Response, []byte, error) {
	startTime := time.Now().UnixNano()
	client := http.Client{
		Timeout: time.Duration(msTimeout) * time.Millisecond,
	}
	if contextType == "" {
		contextType = "application/x-www-form-urlencoded"
	}
	urlParamEncode := urlParams.Encode()
	req, _ := http.NewRequest("POST", urlString, strings.NewReader(urlParamEncode))
	if len(header) > 0 {
		req.Header = header
	}
	req.Header.Set("Content-Type", contextType)
	resp, err := client.Do(req)
	if err != nil {
		logger.Warn("CurlFailed", map[string]interface{}{
			"url":       urlString,
			"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
			"method":    "POST",
			"args":      Substr(urlParamEncode, 0, 1024),
			"err":       err.Error(),
		})
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("CurlFailed", map[string]interface{}{
			"url":       urlString,
			"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
			"method":    "POST",
			"args":      Substr(urlParamEncode, 0, 1024),
			"result":    Substr(string(body), 0, 1024),
			"err":       err.Error(),
		})
		return nil, nil, err
	}
	logger.Info("CurlSuccess", map[string]interface{}{
		"url":       urlString,
		"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		"method":    "POST",
		"args":      Substr(urlParamEncode, 0, 1024),
		"result":    Substr(string(body), 0, 1024),
	})
	return resp, body, nil
}

func CurlJson(urlString string, jsonContent string, msTimeout int, header http.Header) (*http.Response, []byte, error) {
	startTime := time.Now().UnixNano()
	client := http.Client{
		Timeout: time.Duration(msTimeout) * time.Millisecond,
	}
	req, _ := http.NewRequest("POST", urlString, strings.NewReader(jsonContent))
	if len(header) > 0 {
		req.Header = header
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.Warn("CurlFailed", map[string]interface{}{
			"url":       urlString,
			"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
			"method":    "POST",
			"args":      Substr(jsonContent, 0, 1024),
			"err":       err.Error(),
		})
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("CurlFailed", map[string]interface{}{
			"url":       urlString,
			"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
			"method":    "POST",
			"args":      Substr(jsonContent, 0, 1024),
			"result":    Substr(string(body), 0, 1024),
			"err":       err.Error(),
		})
		return nil, nil, err
	}
	logger.Info("CurlSuccess", map[string]interface{}{
		"url":       urlString,
		"proc_time": float32(time.Now().UnixNano()-startTime) / 1.0e9,
		"method":    "POST",
		"args":      Substr(jsonContent, 0, 1024),
		"result":    Substr(string(body), 0, 1024),
	})
	return resp, body, nil
}

func AddGetDataToUrl(urlString string, data url.Values) string {
	if strings.Contains(urlString, "?") {
		urlString = urlString + "&"
	} else {
		urlString = urlString + "?"
	}
	return fmt.Sprintf("%s%s", urlString, data.Encode())
}

//Substr 字符串的截取
func Substr(str string, start int64, end int64) string {
	length := int64(len(str))
	if start < 0 || start > length {
		return ""
	}
	if end < 0 {
		return ""
	}
	if end > length {
		end = length
	}
	return string(str[start:end])
}
