package api

import (
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"github.com/Azunyan1111/multilogin/mysql"
	"log"
	"encoding/json"
	"strconv"
)

func TestGetAge(t *testing.T) {
	mysql.DataBaseInit()
	e, req, rec := myHandler.TestTemplateGet("/api/user/age?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555")
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAge(c)) {
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		if err != nil{
			panic(err)
		}
		var persons AgeResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t,nil,err)
		_, err = strconv.Atoi(persons.Age)
		assert.Equal(t,nil,err)
	}
}
