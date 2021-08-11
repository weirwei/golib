package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type Options struct {
	URL         string
	Body        map[string]string
	JsonBody    string
	ContentType string
	Headers     map[string]string
	Cookies     map[string]string
}

type Result struct {
	HttpCode     int
	ResponseBody []byte
}

func HttpPost(opt *Options) (*Result, error) {
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

func HttpGet(opt *Options) (*Result, error) {
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
	if len(o.Body) > 0 {
		values := url.Values{}
		for k, v := range o.Body {
			values.Set(k, v)
		}
		data = values.Encode()
	} else if len(o.JsonBody) > 0 {
		var err error
		data, err = jsoniter.MarshalToString(o.JsonBody)
		if err != nil {
			return "", err
		}
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
	if len(o.Body) > 0 {
		o.ContentType = contentTypeForm
		return
	}
	if len(o.JsonBody) > 0 {
		o.ContentType = contentTypeJson
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
