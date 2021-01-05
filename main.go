package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)




func main() {
	httpClient()
}

func HttpServer()  {
	mux:=http.DefaultServeMux

	mux.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		key:=request.URL.Query().Get("key")

		writer.Write([]byte(key))
	})

	s:=http.Server{
		Addr:"127.0.0.1:4000",
		Handler:mux,
	}

	s.ListenAndServe()


}

func httpClient()  {
	client:=getHttpClient()

	for i:=0;i<10;i++ {
		go func(i int) {
			url:=fmt.Sprintf("http://127.0.0.1:4000/a?key=%d",i)
			req,_:=http.NewRequest("GET", url, nil)
			if resp, err := client.Do(req); err == nil {
				if bodyByte, err := ioutil.ReadAll(resp.Body); err == nil {
					fmt.Printf("请求：%d,响应:%s\n",i,string(bodyByte))
				}

				resp.Body.Close()
			}

		}(i)
	}

	time.Sleep(100*time.Second)

}


func getHttpClient() *http.Client {
	var (
		cookieJar *cookiejar.Jar
	)
	tr := &http.Transport{
		DisableKeepAlives: true,                                  //禁用长连接
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, //跳过证书验证
	}
	//http cookie接口
	cookieJar, _ = cookiejar.New(nil);

	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	return client

}