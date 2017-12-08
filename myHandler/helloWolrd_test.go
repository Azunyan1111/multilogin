package myHandler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)


func TestHandler_HelloWorld(t *testing.T) {
	e, req, rec := testTemplate("/")
	c := e.NewContext(req, rec)

	if assert.NoError(t, HelloWorld(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello World", rec.Body.String())
	}
}
