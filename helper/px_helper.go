package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	px "qingliuBot/model/4px"
	"strings"
)

func PXAlwaysLive(cookie string) error {
	// curl 'https://order.4px.com/order/my_orders/management/query_order_list' \
	//   -H 'authority: order.4px.com' \
	//   -H 'accept: application/json, text/plain, */*' \
	//   -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
	//   -H 'content-type: application/x-www-form-urlencoded;charset=UTF-8' \
	//   -H 'cookie: Hm_lvt_f718efe499d5223e8942267b94e26693=1660975927; Hm_lvt_9fda412926e52b163fabf7d39d135ad4=1660975931; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODJiOWZjMjUxMzYwNC0wZTk2Y2E2OTdkODFiODgtMjYwMjFkNTEtMzY4NjQwMC0xODJiOWZjMjUxNDk0ZSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%7D; _ga=GA1.1.1267281640.1660980595; _uetvid=e2be2270205911edb420a10e020a76a4; _clck=yh5ig|1|f46|0; _ga_GG67RF9R9K=GS1.1.1660995346.2.0.1660995346.60.0.0; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663395413,1663570122; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1663570122; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663395415,1663570124; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1663570124; tgt_cn=TGT-139627-SdlOPBtjuP2SflTXfOcFbl2EweucLTInKUtwNAPqMvM9pxczkl-www.4px.com; SameSite=None; Auth-Status=loginSuccess; route=1663570135.646.58750.628815; SESSION=NzBhZDhiZTgtYTE5ZC00ZDQzLTk2YzQtM2YwZjNlODk4ZGEy' \
	//   -H 'locale: zh' \
	//   -H 'origin: https://b.4px.com' \
	//   -H 'referer: https://b.4px.com/' \
	//   -H 'sec-ch-ua: "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "Windows"' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-site: same-site' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36' \
	//   -H 'x-requested-with: XMLHttpRequest' \
	//   --data-raw 'pageNum=1&pageSize=20&statusName=INFO_RECEIVED&startDate=2022-09-05%2000%3A00%3A00&endDate=2022-09-19%2023%3A59%3A59&isHistory=N' \
	//   --compressed

	body := strings.NewReader("pageNum=1&pageSize=20&statusName=INFO_RECEIVED&startDate=2022-09-05%2000%3A00%3A00&endDate=2022-09-19%2023%3A59%3A59&isHistory=N")
	req, err := http.NewRequest("POST", "https://order.4px.com/order/my_orders/management/query_order_list", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "order.4px.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	//req.Header.Set("Cookie", "Hm_lvt_f718efe499d5223e8942267b94e26693=1660975927; Hm_lvt_9fda412926e52b163fabf7d39d135ad4=1660975931; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODJiOWZjMjUxMzYwNC0wZTk2Y2E2OTdkODFiODgtMjYwMjFkNTEtMzY4NjQwMC0xODJiOWZjMjUxNDk0ZSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%7D; _ga=GA1.1.1267281640.1660980595; _uetvid=e2be2270205911edb420a10e020a76a4; _clck=yh5ig|1|f46|0; _ga_GG67RF9R9K=GS1.1.1660995346.2.0.1660995346.60.0.0; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663395413,1663570122; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1663570122; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663395415,1663570124; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1663570124; tgt_cn=TGT-139627-SdlOPBtjuP2SflTXfOcFbl2EweucLTInKUtwNAPqMvM9pxczkl-www.4px.com; SameSite=None; Auth-Status=loginSuccess; route=1663570135.646.58750.628815; SESSION=NzBhZDhiZTgtYTE5ZC00ZDQzLTk2YzQtM2YwZjNlODk4ZGEy")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Locale", "zh")
	req.Header.Set("Origin", "https://b.4px.com")
	req.Header.Set("Referer", "https://b.4px.com/")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"105\", \"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"105\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	return nil
}

func Get4PXOrderList(cookie string, startDate string, endDate string, page string, orderNo string) (px.OrderList, error) {
	// curl 'https://order.4px.com/order/my_orders/management/query_order_list' \
	//   -H 'authority: order.4px.com' \
	//   -H 'accept: application/json, text/plain, */*' \
	//   -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
	//   -H 'content-type: application/x-www-form-urlencoded;charset=UTF-8' \
	//   -H 'cookie: Hm_lvt_f718efe499d5223e8942267b94e26693=1660975927; Hm_lvt_9fda412926e52b163fabf7d39d135ad4=1660975931; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODJiOWZjMjUxMzYwNC0wZTk2Y2E2OTdkODFiODgtMjYwMjFkNTEtMzY4NjQwMC0xODJiOWZjMjUxNDk0ZSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%7D; _ga=GA1.1.1267281640.1660980595; _uetvid=e2be2270205911edb420a10e020a76a4; _clck=yh5ig|1|f46|0; _ga_GG67RF9R9K=GS1.1.1660995346.2.0.1660995346.60.0.0; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663395413,1663570122; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1663570122; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663395415,1663570124; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1663570124; tgt_cn=TGT-139627-SdlOPBtjuP2SflTXfOcFbl2EweucLTInKUtwNAPqMvM9pxczkl-www.4px.com; SameSite=None; Auth-Status=loginSuccess; route=1663570135.646.58750.628815; SESSION=NzBhZDhiZTgtYTE5ZC00ZDQzLTk2YzQtM2YwZjNlODk4ZGEy' \
	//   -H 'locale: zh' \
	//   -H 'origin: https://b.4px.com' \
	//   -H 'referer: https://b.4px.com/' \
	//   -H 'sec-ch-ua: "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "Windows"' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-site: same-site' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36' \
	//   -H 'x-requested-with: XMLHttpRequest' \
	//   --data-raw 'pageNum=1&pageSize=20&statusName=INFO_RECEIVED&startDate=2022-09-05%2000%3A00%3A00&endDate=2022-09-19%2023%3A59%3A59&isHistory=N' \
	//   --compressed
	bodyStr := "pageNum=" + page + "&pageSize=500&statusName=INFO_RECEIVED&startDate=" + startDate + "%2000%3A00%3A00&endDate=" + endDate + "%2023%3A59%3A59&isHistory=N"
	//body := strings.NewReader("startDate=2022-09-12%2000%3A00%3A00&endDate=2022-09-26%2023%3A59%3A59&isHistory=N&pageNum=1&pageSize=20&queryOrderNo=XMZX17010943&countryCode=&productCode=&startPrintDate=&endPrintDate=&consigneeName=&subAccountId=&accountType=&startStockInDate=&endStockInDate=&distributorCode=")
	if len(orderNo) > 0 {

		// curl 'https://order.4px.com/order/my_orders/management/query_order_list' \
		//   -H 'authority: order.4px.com' \
		//   -H 'accept: application/json, text/plain, */*' \
		//   -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
		//   -H 'content-type: application/x-www-form-urlencoded;charset=UTF-8' \
		//   -H 'cookie: SameSite=None; SameSite=None; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%2218280417720908-04d5046cb892a58-1c525635-1296000-18280417721815%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODI4MDQxNzcyMDkwOC0wNGQ1MDQ2Y2I4OTJhNTgtMWM1MjU2MzUtMTI5NjAwMC0xODI4MDQxNzcyMTgxNSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%2218280417720908-04d5046cb892a58-1c525635-1296000-18280417721815%22%7D; _clck=a5ikyx|1|f3w|0; _uetvid=e63ac620189511edabc00f7bf39ab976; _ga=GA1.2.1080531220.1660126761; _ga_GG67RF9R9K=GS1.1.1660126760.1.1.1660126776.44; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663324308; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663324309; SameSite=None; Auth-Status=loginSuccess; org.springframework.web.servlet.i18n.CookieLocaleResolver.LOCALE=zh-CN; SERVERID=b43ec71182f0d1026b922efeb4178aa2|1663927283|1663926706; route=1664025239.48.365992.52711; SESSION=ZmQyZjk3ZjQtMTI0OS00NjdiLTg3ZGYtMDQxYmZkMDdhNzI4; tgt_cn=TGT-230016-yNoaKj1oGraeuvr32UPCnWYippjYEvXrywpgd73CIlFyVaJHws-www.4px.com; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1664159341; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1664159342' \
		//   -H 'locale: zh' \
		//   -H 'origin: https://b.4px.com' \
		//   -H 'referer: https://b.4px.com/' \
		//   -H 'sec-ch-ua: "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"' \
		//   -H 'sec-ch-ua-mobile: ?0' \
		//   -H 'sec-ch-ua-platform: "macOS"' \
		//   -H 'sec-fetch-dest: empty' \
		//   -H 'sec-fetch-mode: cors' \
		//   -H 'sec-fetch-site: same-site' \
		//   -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36' \
		//   -H 'x-requested-with: XMLHttpRequest' \
		//   --data-raw 'statusName=INFO_RECEIVED&startDate=2022-09-12%2000%3A00%3A00&endDate=2022-09-26%2023%3A59%3A59&isHistory=N&pageNum=1&pageSize=20&queryOrderNo=XMZX17010943&countryCode=&productCode=&startPrintDate=&endPrintDate=&consigneeName=&subAccountId=&accountType=&startStockInDate=&endStockInDate=&distributorCode=' \
		//   --compressed

		//body := strings.NewReader("statusName=INFO_RECEIVED&startDate=2022-09-12%2000%3A00%3A00&endDate=2022-09-26%2023%3A59%3A59&isHistory=N&pageNum=1&pageSize=20&queryOrderNo=XMZX17010943&countryCode=&productCode=&startPrintDate=&endPrintDate=&consigneeName=&subAccountId=&accountType=&startStockInDate=&endStockInDate=&distributorCode=")

		bodyStr += "&queryOrderNo=" + orderNo + "&countryCode=&productCode=&startPrintDate=&endPrintDate=&consigneeName=&subAccountId=&accountType=&startStockInDate=&endStockInDate=&distributorCode="
	}
	body := strings.NewReader(bodyStr)
	req, err := http.NewRequest("POST", "https://order.4px.com/order/my_orders/management/query_order_list", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "order.4px.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	//req.Header.Set("Cookie", "Hm_lvt_f718efe499d5223e8942267b94e26693=1660975927; Hm_lvt_9fda412926e52b163fabf7d39d135ad4=1660975931; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODJiOWZjMjUxMzYwNC0wZTk2Y2E2OTdkODFiODgtMjYwMjFkNTEtMzY4NjQwMC0xODJiOWZjMjUxNDk0ZSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%7D; _ga=GA1.1.1267281640.1660980595; _uetvid=e2be2270205911edb420a10e020a76a4; _clck=yh5ig|1|f46|0; _ga_GG67RF9R9K=GS1.1.1660995346.2.0.1660995346.60.0.0; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663395413,1663570122; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1663570122; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663395415,1663570124; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1663570124; tgt_cn=TGT-139627-SdlOPBtjuP2SflTXfOcFbl2EweucLTInKUtwNAPqMvM9pxczkl-www.4px.com; SameSite=None; Auth-Status=loginSuccess; route=1663570135.646.58750.628815; SESSION=NzBhZDhiZTgtYTE5ZC00ZDQzLTk2YzQtM2YwZjNlODk4ZGEy")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Locale", "zh")
	req.Header.Set("Origin", "https://b.4px.com")
	req.Header.Set("Referer", "https://b.4px.com/")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"105\", \"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"105\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	orderList := px.OrderList{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return orderList, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	code := resp.StatusCode
	fmt.Println(code)
	if err != nil {
		return orderList, err
	}
	err = json.Unmarshal(data, &orderList)
	if err != nil {
		fmt.Println(string(data))
		fmt.Println(err)
		return orderList, err
	}
	return orderList, nil

}

func GetOrderDetail(cookie, id string) (px.OrderDetail, error) {
	// curl 'https://order.4px.com/order/my_orders/management/query_order_detail' \
	//   -H 'authority: order.4px.com' \
	//   -H 'accept: application/json, text/plain, */*' \
	//   -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
	//   -H 'content-type: application/x-www-form-urlencoded;charset=UTF-8' \
	//   -H 'cookie: SameSite=None; Hm_lvt_f718efe499d5223e8942267b94e26693=1660975927; Hm_lvt_9fda412926e52b163fabf7d39d135ad4=1660975931; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODJiOWZjMjUxMzYwNC0wZTk2Y2E2OTdkODFiODgtMjYwMjFkNTEtMzY4NjQwMC0xODJiOWZjMjUxNDk0ZSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%7D; _ga=GA1.1.1267281640.1660980595; _uetvid=e2be2270205911edb420a10e020a76a4; _clck=yh5ig|1|f46|0; _ga_GG67RF9R9K=GS1.1.1660995346.2.0.1660995346.60.0.0; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663395413,1663570122; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1663570122; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663395415,1663570124; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1663570124; tgt_cn=TGT-139627-SdlOPBtjuP2SflTXfOcFbl2EweucLTInKUtwNAPqMvM9pxczkl-www.4px.com; SameSite=None; Auth-Status=loginSuccess; route=1663570135.646.58750.628815; SESSION=NzBhZDhiZTgtYTE5ZC00ZDQzLTk2YzQtM2YwZjNlODk4ZGEy' \
	//   -H 'locale: zh' \
	//   -H 'origin: https://b.4px.com' \
	//   -H 'referer: https://b.4px.com/' \
	//   -H 'sec-ch-ua: "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'sec-ch-ua-platform: "Windows"' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-site: same-site' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36' \
	//   -H 'x-requested-with: XMLHttpRequest' \
	//   --data-raw 'orderId=8390667612&statusName=INFO_RECEIVED' \
	//   --compressed

	body := strings.NewReader("orderId=" + id + "&statusName=INFO_RECEIVED")
	req, err := http.NewRequest("POST", "https://order.4px.com/order/my_orders/management/query_order_detail", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "order.4px.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	//req.Header.Set("Cookie", "SameSite=None; Hm_lvt_f718efe499d5223e8942267b94e26693=1660975927; Hm_lvt_9fda412926e52b163fabf7d39d135ad4=1660975931; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%224887d7860f71462c86e998d196117815%22%2C%22first_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfbG9naW5faWQiOiI0ODg3ZDc4NjBmNzE0NjJjODZlOTk4ZDE5NjExNzgxNSIsIiRpZGVudGl0eV9jb29raWVfaWQiOiIxODJiOWZjMjUxMzYwNC0wZTk2Y2E2OTdkODFiODgtMjYwMjFkNTEtMzY4NjQwMC0xODJiOWZjMjUxNDk0ZSJ9%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%24identity_login_id%22%2C%22value%22%3A%224887d7860f71462c86e998d196117815%22%7D%2C%22%24device_id%22%3A%22182b9fc2513604-0e96ca697d81b88-26021d51-3686400-182b9fc251494e%22%7D; _ga=GA1.1.1267281640.1660980595; _uetvid=e2be2270205911edb420a10e020a76a4; _clck=yh5ig|1|f46|0; _ga_GG67RF9R9K=GS1.1.1660995346.2.0.1660995346.60.0.0; Hm_lvt_d62f3e8f79208003311c236d688ad4a6=1663395413,1663570122; Hm_lpvt_d62f3e8f79208003311c236d688ad4a6=1663570122; Hm_lvt_00a97f83d9d4c366bd862b5df01a95ee=1663395415,1663570124; Hm_lpvt_00a97f83d9d4c366bd862b5df01a95ee=1663570124; tgt_cn=TGT-139627-SdlOPBtjuP2SflTXfOcFbl2EweucLTInKUtwNAPqMvM9pxczkl-www.4px.com; SameSite=None; Auth-Status=loginSuccess; route=1663570135.646.58750.628815; SESSION=NzBhZDhiZTgtYTE5ZC00ZDQzLTk2YzQtM2YwZjNlODk4ZGEy")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Locale", "zh")
	req.Header.Set("Origin", "https://b.4px.com")
	req.Header.Set("Referer", "https://b.4px.com/")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"105\", \"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"105\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	orderDetail := px.OrderDetail{}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return orderDetail, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orderDetail, err
	}
	err = json.Unmarshal(data, &orderDetail)
	if err != nil {
		fmt.Println(string(data))
		fmt.Println(err)
		return orderDetail, err
	}
	return orderDetail, nil
}
