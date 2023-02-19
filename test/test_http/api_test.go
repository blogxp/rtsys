package test_http

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestApiUser(t *testing.T)  {
	engine := TestLoadEnv()

	w := httptest.NewRecorder()
	//注意第三个参数不能为空，否则会出现 invalid memory address or nil pointer dereference
	r, _ := http.NewRequest(http.MethodGet, "/api/user/public/index?name=奥瑟大撒大", bytes.NewReader([]byte{}))

	engine.ServeHTTP(w,r)

	fmt.Println(w.Body)
}

//post json请求
func TestApiUserPostJson(t *testing.T)  {
	engine := TestLoadEnv()

	w := httptest.NewRecorder()
	params := map[string]any{"name":"国发电商"}
	bytess, _ := json.Marshal(params)
	r, _ := http.NewRequest(http.MethodPost, "/api/user/public/json", bytes.NewBuffer(bytess))
	engine.ServeHTTP(w,r)

	fmt.Println(w.Body)
}

//post  form参数请求
func TestApiUserPostForm(t *testing.T)  {
	engine := TestLoadEnv()

	w := httptest.NewRecorder()

	//由于在api的登录判断中间件中使用了ShouldBindBodyWith，会导致无法通过PostForm获取到

	//post的form参数
	params := url.Values{}
	params.Add("name","123")

	r, _ := http.NewRequest(http.MethodPost, "/api/user/public/form", bytes.NewBuffer([]byte(params.Encode())))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	engine.ServeHTTP(w,r)

	fmt.Println(w.Body)
}