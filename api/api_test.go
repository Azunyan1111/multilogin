package api

import (
	"encoding/json"
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSum(t *testing.T) {
	e, req, rec := myHandler.TestTemplateGet("/api/sum?value1=1&value2=1")
	c := e.NewContext(req, rec)

	if assert.NoError(t, Sum(c)) {
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		if err != nil {
			panic(err)
		}
		var persons SumResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, persons.Value)
	}
}
