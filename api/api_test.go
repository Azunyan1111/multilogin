package api

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/Azunyan1111/multilogin/myHandler"
	"log"
	"github.com/Azunyan1111/multilogin/mysql"
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


func TestGetName(t *testing.T) {
	mysql.DataBaseInit()
	e, req, rec := myHandler.TestTemplateGet("/api/user/name?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555")
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetName(c)) {
		log.Println(rec.Result().StatusCode)
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		if err != nil{
			panic(err)
		}
		log.Println(string(byteArray))
		var persons NameResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t,nil,err)
		assert.Equal(t,"Azunyan1111",persons.Name)
	}
}