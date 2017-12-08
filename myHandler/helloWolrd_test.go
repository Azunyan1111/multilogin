package myHandler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"github.com/PuerkitoBio/goquery"
)


func TestHelloWorld(t *testing.T) {
	e, req, rec := testTemplateGet("/")
	c := e.NewContext(req, rec)

	if assert.NoError(t, HelloWorld(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_helloWorld").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "Hello World",text)
	}
}
