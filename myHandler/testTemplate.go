package myHandler

import (
	"github.com/Azunyan1111/multilogin/model"
	"github.com/Azunyan1111/multilogin/mysql"
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"
)

const (
	userUid2    = "26d2983e-3d5a-421c-bf6f-d4608025e555"
	serviceUid2 = "124ah368-1eha-7h81-2345-365a24h6522y"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TestTemplateGet(target string) (e *echo.Echo, req *http.Request, rec *httptest.ResponseRecorder) {
	e = echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	temp := &Template{
		templates: template.Must(template.ParseGlob("../static/views/*.html")),
	}
	e.Renderer = temp

	req = httptest.NewRequest(echo.GET, target, nil)
	rec = httptest.NewRecorder()

	e, req, rec = setAuthorization(e, req, rec, target)

	return e, req, rec
}
func TestTemplatePost(target string, json string) (e *echo.Echo, req *http.Request, rec *httptest.ResponseRecorder) {
	e = echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	temp := &Template{
		templates: template.Must(template.ParseGlob("../static/views/*.html")),
	}
	e.Renderer = temp

	req = httptest.NewRequest(echo.POST, target, strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()

	e, req, rec = setAuthorization(e, req, rec, target)

	return e, req, rec
}

func setAuthorization(e *echo.Echo, req *http.Request, rec *httptest.ResponseRecorder, url string) (*echo.Echo, *http.Request, *httptest.ResponseRecorder) {
	// ヘッダーに付ける奴
	var authorization string
	authorization += "MLAuth1.0" + ","

	// サービス情報を取得
	orm := mysql.GetOrm()
	var service structs.Service
	orm.Find(&service, "uuid = ?", serviceUid2)
	authorization += "token=" + service.Token + ","

	// リクエスト情報をカンマでつなげる
	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	authorization += "time=" + timeStr + ","

	// シークレットトークンから繋げたstringでハッシュを入手
	join := url + "," + timeStr
	hash := model.GetHmac(service.Secret, join)
	authorization += "signature=" + hash + ","
	req.Header.Set("authorization", authorization)

	return e, req, rec
}
