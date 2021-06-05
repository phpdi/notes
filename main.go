package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

/*

占位符是${},解析字符串，比如
1.${a}=${b} 解析后得出a=b
2.${a}=${b\} 解析后得出a=${b}
3.${a\}=${b} 解析后得出${a}=b
4.${a\}=${b\} 解析后得出${a}=${b}
5.$${a}=${b} 解析后得出$a=b
6.$${a}=${b}} 解析后得出$a=b}
要求，不用正则表达式，不准用split之类的函数，请遍历字符串。

*/

func main() {
	c := make(chan bool)
	errgo()

	<-c
}

func errgo() {

	lock := new(sync.Mutex)
	a := make([]int, 1000)

	l := len(a)
	num := 0
	for k := range a {
		a[k] = 1
	}

	for _, v := range a {
		v := v
		go func() {
			defer func() {
				lock.Lock()
				num += v
				if num == l {
					fmt.Println(num)
				}

				lock.Unlock()

			}()
		}()
	}

}

func fmtEnv() {
	envParams := `PROFILE=KS;SERVICE=pricing-manager;TZ=UTC;BEATFLOW_ENABLED_FEATURE_IGNORE_AUTHORITIES=true;BEATFLOW_KS_MYSQL_HOST=103.61.39.4;BEATFLOW_KS_MYSQL_PORT=13306;BEATFLOW_KS_MYSQL_DATABASE=pricing;BEATFLOW_KS_MYSQL_PASSWORD=p@ss52Dnb;BEATFLOW_DEV_MYSQL_HOST=103.61.39.4;BEATFLOW_DEV_MYSQL_PORT=23306;BEATFLOW_DEV_MYSQL_DATABASE=pricing;BEATFLOW_DEV_MYSQL_PASSWORD=p@ss52Dnb;BEATFLOW_KS_REDIS_ENABLE=true;BEATFLOW_KS_REDIS_URL=103.61.39.4:16000;BEATFLOW_DEV_REDIS_ENABLE=true;BEATFLOW_DEV_REDIS_URL=103.61.39.4:26839`
	s := strings.Split(envParams, ";")
	for _, v := range s {
		ss := strings.Split(v, "=")
		fmt.Printf("export %s=%s \n", ss[0], ss[1])
	}

}

func parse(in string) (out string) {

	l := len(in)
A:
	for i := 0; i < l; i++ {
		next := i + 1
		if string(in[i]) == "$" && string(in[next]) == "{" {
			end := next + 1
			for end < l {
				if string(in[end]) == "}" {

					//转义情况
					if string(in[end-1]) == "\\" {
						out += in[i:end-1] + "}"
					} else {
						out += in[next+1 : end]
					}

					i = end
					continue A
				}
				end++
			}
		}

		out += in[i:next]

	}

	return
}

func writer() {
	w := tabwriter.NewWriter(os.Stderr, 3, 0, 3, ' ', tabwriter.TabIndent|tabwriter.Debug)

	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // no trailing tab
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	w.Flush()
}

func HttpServer() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		key := request.URL.Query().Get("key")

		writer.Write([]byte(key))
	})

	s := http.Server{
		Addr:    "127.0.0.1:4000",
		Handler: mux,
	}

	s.ListenAndServe()

}

func httpClient() {
	client := getHttpClient()

	for i := 0; i < 10; i++ {
		go func(i int) {
			url := fmt.Sprintf("http://127.0.0.1:4000/a?key=%d", i)
			req, _ := http.NewRequest("GET", url, nil)
			if resp, err := client.Do(req); err == nil {
				if bodyByte, err := ioutil.ReadAll(resp.Body); err == nil {
					fmt.Printf("请求：%d,响应:%s\n", i, string(bodyByte))
				}

				resp.Body.Close()
			}

		}(i)
	}

	time.Sleep(100 * time.Second)

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
	cookieJar, _ = cookiejar.New(nil)

	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	return client

}
