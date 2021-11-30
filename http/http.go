package http

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

const (
	// EncodeJson 请求数据类型为json
	EncodeJson = "_json"

	// EncodeForm 请求数据类型为form
	EncodeForm = "_form"
)

// Options http request options
// URL request url
// RequestBody 请求体
// Encode default form
// Headers headers
// Cookies cookies
type Options struct {
	URL         string
	RequestBody interface{}
	ContentType string
	Encode      string
	Headers     map[string]string
	Cookies     map[string]string
}

// Result http request result
type Result struct {
	HttpCode     int
	ResponseBody []byte
}

// Post http post request
func Post(opt *Options) (*Result, error) {
	data, err := opt.getData()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(httpPost, opt.URL, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	opt.makeRequest(request)
	client := http.Client{}
	log.Printf("post request:%v", opt)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	log.Printf("post response:%v", response)
	res, err := responseToResult(response)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Get(opt *Options) (*Result, error) {
	request, err := http.NewRequest(httpGet, opt.URL, nil)
	client := http.Client{}
	log.Printf("get request:%v", opt)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	log.Printf("get response:%v", response)
	res, err := responseToResult(response)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (o *Options) getData() (string, error) {
	var data string
	var err error
	switch o.Encode {
	case EncodeJson:
		data, err = jsoniter.MarshalToString(o.RequestBody)
		if err != nil {
			return "", err
		}
	case EncodeForm:
		fallthrough
	default:
		value := &url.Values{}
		if formData, ok := o.RequestBody.(map[string]string); ok {
			for k, v := range formData {
				value.Set(k, v)
			}
		} else if formData, ok := o.RequestBody.(map[string]interface{}); ok {
			for k, v := range formData {
				switch v.(type) {
				case string:
					value.Set(k, v.(string))
				default:
					vStr, err := jsoniter.MarshalToString(v)
					if err != nil {
						return data, err
					}
					value.Set(k, vStr)
				}
			}
		} else {
			return data, errors.New("get requestBody error")
		}
		data = value.Encode()
	}

	return data, nil
}

func (o *Options) makeRequest(req *http.Request) {
	for key, val := range o.Headers {
		req.Header.Set(key, val)
	}
	o.getContentType()
	req.Header.Set("Content-Type", o.ContentType)
	for key, val := range o.Cookies {
		req.AddCookie(&http.Cookie{
			Name:  key,
			Value: val,
		})
	}
}

func (o *Options) getContentType() {
	if len(o.ContentType) != 0 {
		return
	}
	switch o.Encode {
	case EncodeJson:
		o.ContentType = contentTypeJson
	case EncodeForm:
		fallthrough
	default:
		o.ContentType = contentTypeForm
	}
}

func responseToResult(response *http.Response) (*Result, error) {
	var res Result
	if response != nil {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		res.HttpCode = response.StatusCode
		res.ResponseBody = body
		_ = response.Body.Close()
	}
	return &res, nil
}
