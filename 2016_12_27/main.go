// 使用golang实现https服务
//
// 证书可使用certbot-auto生成
//      步骤可参照：https://certbot.eff.org/#ubuntutrusty-nginx
//
package main

import (
	"net/http"
)

type httpHandler struct {
}

func (a httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello i am http."))
}

func listenAndServeHttp() {
	port := ":5920"
	http.Handle("/", httpHandler{})
	err := http.ListenAndServe(port, nil)
	panic(err)
}

func main() {
	// listenAndServeHttp()

	listenAndServeHttps()
}

// TODO
func listenAndServeHttps() {

}
