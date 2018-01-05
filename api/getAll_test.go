package api

import (
	"encoding/json"
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestGetAll(t *testing.T) {
	mysql.DataBaseInit()
	e, req, rec := myHandler.TestTemplateGet("/api/user/all?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555")
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAll(c)) {
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		if err != nil {
			panic(err)
		}
		var persons AllResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t, nil, err)
		_, err = strconv.Atoi(persons.User.Age)
		assert.Equal(t, "Azunyan1111", persons.User.UserName)
		assert.Equal(t, "azunyan1111@azunyan.me", persons.User.Email)
		assert.Equal(t, "http://noimage.com/azunyan", persons.User.Image)
		assert.Equal(t, nil, err)
		assert.Equal(t, "2017-08-01", persons.User.Birthday)
		assert.Equal(t, "090-1145-1419", persons.User.Phone)
		assert.Equal(t, "福岡", persons.User.Address)

	}
}
