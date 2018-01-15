package api

import (
	"encoding/json"
	"github.com/Azunyan1111/multilogin/myHandler"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGetName(t *testing.T) {
	mysql.DataBaseInit()
	e, req, rec := myHandler.TestTemplateGet("/api/user/name?uuid=26d2983e-3d5a-421c-bf6f-d4608025e555")
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetName(c)) {
		byteArray, err := ioutil.ReadAll(rec.Result().Body)
		assert.NoError(t, err)
		var persons NameResponse
		err = json.Unmarshal(byteArray, &persons)
		assert.Equal(t, nil, err)
		assert.Equal(t, "Azunyan1111", persons.Name)
	}
}
