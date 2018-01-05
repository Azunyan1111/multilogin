package api

import (
	"encoding/json"
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func TestGetPhone(t *testing.T) {
	mysql.DataBaseInit()
	e, req, rec := myHandler.TestTemplateGet("/api/user/phone?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555")
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetPhone(c)) {
		log.Println(rec.Result().StatusCode)
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		if err != nil {
			panic(err)
		}
		var persons PhoneResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t, nil, err)
		assert.Equal(t, "090-1145-1419", persons.Phone)
	}
}
