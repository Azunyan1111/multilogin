package api

import (
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"github.com/Azunyan1111/multilogin/mysql"
	"log"
	"encoding/json"
)

func TestGetImage(t *testing.T) {
	mysql.DataBaseInit()
	e, req, rec := myHandler.TestTemplateGet("/api/user/image?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555")
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetImage(c)) {
		log.Println(rec.Result().StatusCode)
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		if err != nil{
			panic(err)
		}
		log.Println(string(byteArray))
		var persons ImageResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t,nil,err)
		assert.Equal(t,"http://noimage.com/azunyan",persons.Image)
	}
}