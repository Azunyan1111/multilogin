package api

import (
	"encoding/json"
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	userUid    = "26d2983e-3d5a-421c-bf6f-d4608025e555"
	serviceUid = "124ah368-1eha-7h81-2345-365a24h6522y"
)

func TestSum(t *testing.T) {
	mysql.DataBaseInit()
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
