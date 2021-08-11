package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpPost(t *testing.T) {
	jsonStr := "{\"category\":\"trending\"," +
		"\"period\":\"day\"," +
		"\"lang\":\"go\"," +
		"\"offset\": 0," +
		"\"limit\": 2}"
	options := Options{
		URL:      "https://e.juejin.cn/resources/github",
		JsonBody: jsonStr,
	}
	result, err := HttpPost(&options)
	assert.Nil(t, err)
	t.Log(string(result.ResponseBody))
}

func TestHttpGet(t *testing.T) {
	options := Options{
		URL: "http://baidu.com",
	}
	result, err := HttpGet(&options)
	assert.Nil(t, err)
	t.Log(string(result.ResponseBody))
}
