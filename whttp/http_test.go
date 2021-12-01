package whttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpPost(t *testing.T) {
	requestMap := map[string]interface{}{
		"category": "treading",
		"period":   "day",
		"lang":     "go",
		"offset":   0,
		"limit":    2,
	}
	options := Options{
		URL:         "https://e.juejin.cn/resources/github",
		RequestBody: requestMap,
		Encode:      EncodeForm,
	}
	result, err := Post(&options)
	assert.Nil(t, err)
	t.Log(string(result.ResponseBody))
}

func TestHttpGet(t *testing.T) {
	requestMap := map[string]interface{}{
		"code": "utf-8",
		"q":    "ps5",
	}
	options := Options{
		URL:         "https://suggest.taobao.com/sug",
		RequestBody: requestMap,
	}
	result, err := Get(&options)
	assert.Nil(t, err)
	t.Log(string(result.ResponseBody))
}
