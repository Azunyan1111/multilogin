package api

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/Azunyan1111/multilogin/myHandler"
)

func TestSum(t *testing.T) {
	e, req, rec := myHandler.TestTemplateGet("/api/sum?value1=1&value2=1")
	c := e.NewContext(req, rec)

	if assert.NoError(t, Sum(c)) {
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		if err != nil{
			panic(err)
		}
		var persons SumResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t,nil,err)
		assert.Equal(t,2,persons.Value)
	}
}