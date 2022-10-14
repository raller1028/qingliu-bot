/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-08-20 14:21:36
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-08-20 14:56:12
 * @FilePath: /feishu-bot/service/4px.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"qingliuBot/helper"
	px "qingliuBot/model/4px"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

type Px4Service interface {
	Login()
	Forward(method string, v string, body []byte) (string, error)
	GetOrderList(orderNo string, startDate, endDate string) (px.OrderList, error)
	GetOrderListYesterday() (px.OrderList, error)
	GetOrderDetail(id string) (px.OrderDetail, error)
}

type px4Service struct {
}

// 登陆客户端
var Client http.Client

func (p *px4Service) Login() {
	req, _ := http.NewRequest("GET", "https://sso.4px.com/login?service=https://b.4px.com/cas&amp;from=bcp&locale=zh", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")

	// 初始化client
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	Client = http.Client{Jar: jar}
	resp, _ := Client.Do(req)

	dom, _ := goquery.NewDocumentFromReader(resp.Body)
	val, _ := dom.Find("input[name=execution]").Attr("value")
	key, _ := dom.Find("#key").Attr("value")
	// return val

	// 构造请求
	param := url.Values{}
	param.Set("username", "13818662085")
	param.Set("password", "Icewhale2020")
	param.Set("execution", val)
	param.Set("lt", "")
	param.Set("imgverifycode", "")
	param.Set("key", key)

	data := param.Encode()
	reqLogin, _ := http.NewRequest("POST", "https://sso.4px.com/login?service=https://b.4px.com/cas", strings.NewReader(data))
	//reqLogin.Header.Set("authority", "sso.4px.com")
	reqLogin.Header.Set("content-type", "application/x-www-form-urlencoded")
	reqLogin.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36")
	reqLogin.Header.Set("referer", "https://sso.4px.com")

	// 发起请求
	respLogin, _ := Client.Do(reqLogin)
	source, _ := ioutil.ReadAll(respLogin.Body)
	fmt.Println(string(source))

}

func (p *px4Service) Forward(method string, v string, body []byte) (string, error) {
	data := make(map[string]string)
	order := []string{"app_key", "format", "method", "timestamp", "v"}
	data["app_key"] = "7a57dc0a-98fa-48e3-ae12-b06cbaf458b8"
	data["format"] = "json"
	data["method"] = method
	data["timestamp"] = fmt.Sprint(time.Now().UnixNano() / 1e6)
	data["v"] = v
	appSecret := "87f8d870-f464-44bd-8a98-80eb4208792b"
	var str = ""
	for _, v := range order {
		str += v + data[v]
	}
	fmt.Println("body:", string(body))
	str += string(body) + appSecret
	md5Data := []byte(str)
	fmt.Println("str:", str)
	has := md5.Sum(md5Data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	fmt.Println("md5", md5str1)
	result := make(map[string]string)
	result["sign"] = md5str1
	result["timestamp"] = data["timestamp"]

	query := ""
	for _, v := range order {
		query += v + "=" + data[v] + "&"
	}
	query += "sign=" + md5str1
	resp, err := http.Post("https://open.4px.com/router/api/service?"+query, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rda, _ := ioutil.ReadAll(resp.Body)
	return string(rda), nil
}
func (p *px4Service) GetOrderList(orderNo string, startDate, endDate string) (px.OrderList, error) {
	return helper.Get4PXOrderList(PXCOOKIE, startDate, endDate, "1", orderNo)
}

func (p *px4Service) GetOrderListYesterday() (px.OrderList, error) {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	startTime := yesTime.Format("2006-01-02")
	fmt.Println("startTime:", startTime)
	return helper.Get4PXOrderList(PXCOOKIE, startTime, startTime, "1", "")
}

func (p *px4Service) GetOrderDetail(id string) (px.OrderDetail, error) {
	return helper.GetOrderDetail(PXCOOKIE, id)
}
func NewPx4Service() Px4Service {
	return &px4Service{}
}
