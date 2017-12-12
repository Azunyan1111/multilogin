package myHandler

import (
	"github.com/stretchr/testify/assert"
	"github.com/PuerkitoBio/goquery"
	"testing"
	"net/http"
	"github.com/ipfans/echo-session"
)

func TestGetUserMyPage(t *testing.T) {
	e, req, rec := testTemplateGet("/user/mypage")

	s := session.Default(e.AcquireContext())
	if s != nil{
		s.Set("uid", "uuid")
		s.Save()
	}
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetUserMyPage(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(rec.Result().Body)
		var text string
		doc.Find("#test_UserMyPage").Each(func(_ int, s *goquery.Selection) {
			text = s.Text()
		})
		assert.Equal(t, "登録情報更新", text)
	}
}