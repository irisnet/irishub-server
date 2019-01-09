package http

import (
	"crypto/tls"
	"fmt"
	"github.com/irisnet/irishub-server/modules/logger"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var client http.Client
var timeout = 5 * time.Second
var keepAlive = 30 * time.Second

func init() {
	client = http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
			}).DialContext,
		},
	}
}

func Get(url string, body io.Reader) (resp Response) {
	logger.Info.Println(fmt.Sprintf("request uri:%s", url))
	req, err := http.NewRequest(http.MethodGet, url, body)
	if err != nil {
		logger.Error.Printf("NewRequest error:%s", err.Error())
		resp.Code = http.StatusBadRequest
		resp.Error = err
		return
	}
	return do(req)
}

func Post(url, contentType string, body io.Reader) (resp Response) {
	logger.Info.Println(fmt.Sprintf("request uri:%s", url))
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		logger.Error.Printf("NewRequest error:%s", err.Error())
		resp.Code = http.StatusBadRequest
		resp.Error = err
		return
	}
	req.Header.Set("Content-Type", contentType)
	return do(req)
}

func do(req *http.Request) (resp Response) {
	res, err := client.Do(req)
	if err != nil {
		logger.Error.Printf("client.Do:%s", err.Error())
		resp.Code = http.StatusBadRequest
		resp.Error = err
		return
	}

	bz, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		logger.Error.Printf("ioutil.ReadAll:%s", err.Error())
		resp.Code = http.StatusBadRequest
		resp.Error = err
		return
	}

	logger.Info.Println(fmt.Sprintf("http response data:%s", string(bz)))
	resp.Code = res.StatusCode
	resp.Data = bz
	return resp
}

type Response struct {
	Code  int
	Data  []byte
	Error error
}
