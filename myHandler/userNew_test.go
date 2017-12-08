package myHandler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGetUserNew(t *testing.T) {
	e, req, rec := testTemplateGet("/user/new")
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetUserNew(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_userNew").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "ユーザー登録", text)
	}
}

func TestPostUserNew(t *testing.T) {

	f := make(url.Values)
	f.Set("InputEmail", "bar@bar.com")
	f.Set("InputUserName", "bar")
	f.Set("InputImage", "https://upload.wikimedia.org/wikipedia/commons/thumb/4/43/Bar-P1030319.jpg/1200px-Bar-P1030319.jpg")
	f.Set("InputAge", "22")
	f.Set("InputBirthday", "1985-5-3")
	f.Set("InputPhone", "080-3749-7392")
	f.Set("InputAddress", "福岡県")

	e, req, rec := testTemplatePost("/user/new", f.Encode())
	c := e.NewContext(req, rec)

	if assert.NoError(t, PostUserNew(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_userNewPost").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "登録完了", text)
	}
}
