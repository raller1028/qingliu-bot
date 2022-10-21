/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-03-28 14:19:12
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-08-09 14:24:03
 * @FilePath: /feishu-bot/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	// "qingliuBot/opkg/mysql"

	"qingliuBot/route"
	"qingliuBot/service"
	"time"

	"github.com/robfig/cron"
)

func init() {
	service.InitLark()

	// service.Service = service.NewService(mysql.GetDb())
	service.Service = service.NewService()
}

func main() {
	// service.PXCOOKIE = "SameSite=None; SameSite=None; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%2218280417720908-04d5046cb892a58-1c525635-1296000-18280417721815%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODI4MDQxNzcyMDkwOC0wNGQ1MDQ2Y2I4OTJhNTgtMWM1MjU2MzUtMTI5NjAwMC0xODI4MDQxNzcyMTgxNSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%2218280417720908-04d5046cb892a58-1c525635-1296000-18280417721815%22%7D; _clck=a5ikyx|1|f3w|0; _uetvid=e63ac620189511edabc00f7bf39ab976; _ga=GA1.2.1080531220.1660126761; _ga_GG67RF9R9K=GS1.1.1660126760.1.1.1660126776.44; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663324308; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663324309; SameSite=None; Auth-Status=loginSuccess; org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=zh-CN; SERVERID=b43ec71182f0d1026b922efeb4178aa2|1663927283|1663926706; route=1664025239.48.365992.52711; tgt_cn=TGT-230016-yNoaKj1oGraeuvr32UPCnWYippjYEvXrywpgd73CIlFyVaJHws-www.4px.com; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1664159341; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1664159342; SESSION=YzcxNTQwYTUtNjIzYS00YjlkLWJmYTAtOGIwM2U5ZDM1NDBk"
	// service.PXCOOKIE = "SameSite=None; Hm_lvt_f718efe499d5223e8942267b94e26693=1665372729; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%221836471a6271cf-02ec79a4d65a0a-26021c51-2073600-1836471a628137b%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODM2NDcxYTYyNzFjZi0wMmVjNzlhNGQ2NWEwYS0yNjAyMWM1MS0yMDczNjAwLTE4MzY0NzFhNjI4MTM3YiJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%221836471a6271cf-02ec79a4d65a0a-26021c51-2073600-1836471a628137b%22%7D; tgt_cn=TGT-163723-vAhjUVjk6FHNIUjvZYYGwzEHI4kUa3wGNagzVfcbKiMW0OxuFG-www.4px.com; SameSite=None; Auth-Status=loginSuccess; route=1665563389.289.59287.898962; SESSION=NDNlZjQ2MDItNDdkMC00OTFlLWFhNGQtNmMzZDhlOTE1ODJj"
	// service.PXCOOKIE = "sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%22183dcbe35a85a6-09acc60ebfbc938-26021f51-1395396-183dcbe35a97c1%22%2C%22props%22%3A%7B%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTgzZGNiZTM1YTg1YTYtMDlhY2M2MGViZmJjOTM4LTI2MDIxZjUxLTEzOTUzOTYtMTgzZGNiZTM1YTk3YzEiLCIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%22183dcbe35a85a6-09acc60ebfbc938-26021f51-1395396-183dcbe35a97c1%22%7D;tgt_cn=TGT-14609-UGMGSFYz5wX5q2ZzbnJHWJnc0NjfplUGLxdjS2PbEegzvVYjjI-www.4px.com;lang=zh;Hm_lpvt_74c3bc84b03814856741a23f2f296480=1665855797;sajssdk_2015_cross_new_user=1;SESSION=NDZjMjAxZTAtOTg2ZC00NGVhLTlhNzEtMjQwYmMyZGQ1OWY3;"
	cron2 := cron.New()
	// //every day execution
	// err := cron2.AddFunc("0 0/5 * * * *", func() {
	// 	if len(service.PXCOOKIE) == 0 {
	// 		a, b, c := service.LarkCli.Message.Send().ToChatID("oc_3b7ab4655bd6de673cf30cce8ce58b19").SendText(context.Background(), "4px需要设置cookie")
	// 		fmt.Println(a, b, c)
	// 	} else {
	// 		//保活4pxcookie
	// 		test := helper.PXAlwaysLive(service.PXCOOKIE)
	// 		fmt.Println(test)
	// 	}
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//every day execution
	//err = cron2.AddFunc("0 0/1 18 * * *", func() {

	err := cron2.AddFunc("0 0 1 * * *", func() { service.Service.Px4().ImportDataByParam("", "", "") })
	if err != nil {
		fmt.Println(err)
	}
	cron2.Start()

	defer cron2.Stop()

	r := route.InitRoute()

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
