#### go httpClient 出现大量ESTABLISHED ,TIME_WAIT
```go
func (this *Tools) ApiPost(url string, postParams lwyserver.Request) (bodyByte []byte, err error) {

	var (
		postStr   []byte
		cookieJar *cookiejar.Jar
		req       *http.Request
		resp      *http.Response
	)
	if this.Debug {
		fmt.Println(url)
	}

	if postStr, err = json.Marshal(postParams); err != nil {
		return
	}


	tr := &http.Transport{
		DisableKeepAlives: true,//禁用长连接
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},//跳过证书验证
	}
	//http cookie接口
	if cookieJar, err = cookiejar.New(nil); err != nil {
		return
	}

	client := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	if req, err = http.NewRequest("POST", url, strings.NewReader(string(postStr))); err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if resp, err = client.Do(req); err != nil {
		return
	}
	resp.Header.Set("Connection", "close")
	defer resp.Body.Close()

	if bodyByte, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if this.Debug {
		fmt.Println("响应数据:", string(bodyByte))
	}

	return
}

```
* DisableKeepAlives
    >这个参数需要设置为true 否则resp.Body.Close() 并不能起到关闭连接的作用，当HTTP请求完成并且接收到响应后，如果对端的HTTP服务器没有关闭连接，那么这个连接会一直处于ESTABLISHED状态。
    >DisableKeepAlives 设置会true的作用实际上会执行	req.Header.Set("Connection", "close"),告诉服务端本次调用为短连接。

* resp.Header.Set("Connection", "close")
    > 如何使用这句，netstat -an 会有少量TIME_WAIT 
    
* 总结：golang http客户端在发送http请求的时候，需要头信息中声明本次使用的是http短链接。 
