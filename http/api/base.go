package api

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ApiRequest struct {
	url     string
	client  *http.Client
	request *http.Request
	method  string
	params  []string
	forms   []string
}

func NewGetApi(_url string) *ApiRequest {
	api := ApiRequest{url: _url}
	api.params = make([]string, 0)
	api.forms = make([]string, 0)
	api.method = "GET"
	return &api
}

func NewPostApi(_url string) *ApiRequest {
	api := ApiRequest{url: _url}
	api.params = make([]string, 0)
	api.forms = make([]string, 0)
	api.method = "POST"
	return &api
}

//添加get参数
func (api *ApiRequest) AddParam(key string, val string) {
	api.params = append(api.params, url.QueryEscape(key)+"="+url.QueryEscape(val))
}

//添加get参数数组格式
func (api *ApiRequest) AddParamByMap(param map[string]string) {
	for key, val := range param {
		api.params = append(api.params, url.QueryEscape(key)+"="+url.QueryEscape(val))
	}
}

//添加post参数
func (api *ApiRequest) AddForm(key string, val string) {
	api.forms = append(api.forms, key+"="+val)
}

//添加post参数数组格式
func (api *ApiRequest) AddFormByMap(param map[string]string) {
	for key, val := range param {
		api.forms = append(api.forms, key+"="+val)
	}
}

//添加非登录态user+password
func (api *ApiRequest) AddUser(user string) {
	api.url = user + "@" + api.url
}

//添加协议头
func (api *ApiRequest) AddProtocol(protocol string) {
	api.url = protocol + "://" + api.url
}

//添加header信息（url编码）
func (api *ApiRequest) AddHeader(header string, val string) {
	api.request.Header.Add(url.QueryEscape(header), url.QueryEscape(val))
}

//添加header信息（不做url编码）
func (api *ApiRequest) AddHeaderWithoutEncode(header string, val string) {
	api.request.Header.Add(header, val)
}

//添加cookie
func (api *ApiRequest) AddCookie(c *http.Cookie) {
	api.request.AddCookie(c)
}

//初始化client、request
func (api *ApiRequest) InitRequest() bool {
	if len(api.params) > 0 {
		paramStr := strings.Join(api.params, "&")
		api.url += "?" + paramStr
	}
	var err error
	api.client = &http.Client{}
	if api.method == "GET" {
		api.request, err = http.NewRequest("GET", api.url, nil)
	} else if api.method == "POST" {
		api.request, err = http.NewRequest("POST", api.url, strings.NewReader(strings.Join(api.forms, "&")))
	} else {
		err = errors.New("invalid request method")
	}

	if err != nil {
		log.Fatal("ApiBaseLogTag", "init get request err:"+err.Error())
		return false
	}
	return true
}

func (api *ApiRequest) InitPostJsonRequest(json string) bool {
	if len(api.params) > 0 {
		paramStr := strings.Join(api.params, "&")
		api.url += "?" + paramStr
	}
	var err error
	api.client = &http.Client{}
	if api.method == "GET" {
		return false
	} else if api.method == "POST" {
		api.request, err = http.NewRequest("POST", api.url, bytes.NewBuffer([]byte(json)))
	} else {
		err = errors.New("invalid request method")
	}
	if err != nil {
		log.Fatal("ApiBaseLogTag", "init InitPostJsonRequest err:"+err.Error())
		return false
	}
	return true
}

func (api *ApiRequest) GetResponse() []byte {
	response, err := api.client.Do(api.request)
	if err != nil {
		log.Fatal("ApiBaseLogTag", "do response err:"+err.Error())
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("ApiBaseLogTag", "get response body err:"+err.Error())
		return nil
	}
	return body
}
