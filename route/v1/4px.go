/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-08-09 11:26:40
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-08-20 14:29:45
 * @FilePath: /feishu-bot/route/v1/4px.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"qingliuBot/helper"
	"qingliuBot/model/qingliu"
	"qingliuBot/service"
	"strconv"
	"strings"
	"time"
)

func GetSign(w http.ResponseWriter, r *http.Request) {

	//io.Reader
	data := make(map[string]string)
	order := []string{"app_key", "format", "method", "timestamp", "v"}
	data["app_key"] = r.URL.Query().Get("app_key")
	data["format"] = "json"
	data["method"] = r.URL.Query().Get("method")
	data["timestamp"] = fmt.Sprint(time.Now().UnixNano() / 1e6)
	data["v"] = r.URL.Query().Get("v")
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(err)
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
		return
	}
	defer resp.Body.Close()
	rda, _ := ioutil.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(json.RawMessage(string(rda)))
}
func Login(w http.ResponseWriter, r *http.Request) {
	service.Service.Px4().Login()

}

func Import(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//rs, _ := service.Service.Px4().Forward("ds.xms.order.get", "1.1.0", []byte(`{"start_time_of_create_consignment":1662964098000,"end_time_of_create_consignment":1663396098000}`))
	//fmt.Println(service.Service.Px4().Forward("fu.wms.inbound.getlist", "1.0.0", []byte(`{"create_time_start":1662964098000,"create_time_end":1663396098000}`)))
	// orders := px.OrderList{}
	// json.Unmarshal([]byte(rs), &orders)
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	// rs := ""
	order := r.URL.Query().Get("order")
	//fmt.Println(service.Service.Px4().Forward("fu.wms.outbound.getlist", "1.0.0", []byte(`{"create_time_start": 1662964098000,"create_time_end":1663396098000}`)))
	//fmt.Println(service.Service.Px4().Forward("fu.wms.inbound.getlist", "1.0.0", []byte(`{"create_time_start": 1662964098000,"create_time_end":1663396098000}`)))
	list, err := service.Service.Px4().GetOrderList(order, startDate, endDate)
	if err != nil {
		json.NewEncoder(w).Encode(json.RawMessage(string(`"code": 500, "msg": ` + err.Error() + `}`)))
		return
	}

	for _, v := range list.List {
		orderDetail, err := service.Service.Px4().GetOrderDetail(strconv.FormatInt(v.ID, 10))
		if err != nil {
			json.NewEncoder(w).Encode(json.RawMessage(string(`"code": 500, "msg": ` + err.Error() + `}`)))
			return
		}
		data := qingliu.QSource{
			Order:     orderDetail.Order.CustomerOrderNo,
			Server:    orderDetail.Order.CoFpxTrackNo,
			Applicant: "897ce02508e3@lark.qingflow.com",
			Name:      orderDetail.ShipperConsignee.ConsigneeFirstName + orderDetail.ShipperConsignee.ConsigneeLastName,
			Phone:     orderDetail.ShipperConsignee.ConsigneeTelephone,
			Company:   orderDetail.ShipperConsignee.ConsigneeCompany,
			Country:   orderDetail.ShipperConsignee.ConsigneeCountry,
			Email:     orderDetail.ShipperConsignee.ConsigneeMail,
			Province:  orderDetail.ShipperConsignee.ConsigneeProvince,
			City:      orderDetail.ShipperConsignee.ConsigneeCity,
			Address:   orderDetail.ShipperConsignee.ConsigneeAddress1 + "," + orderDetail.ShipperConsignee.ConsigneeAddress2,
			ZipCode:   orderDetail.ShipperConsignee.ConsigneePostcode,
			Trial:     v.TrialAmount,
		}
		if len(data.Name) == 0 {
			data.Name = orderDetail.ShipperConsignee.ConsigneeName
		}
		products := []qingliu.Product{}

		for _, p := range orderDetail.DeclareList {
			skuArr := strings.Split(p.Ename, "(")
			sku := ""
			if len(skuArr) > 1 {
				sku = strings.Split(skuArr[1], ")")[0]
			} else {
				a, b, c := service.LarkCli.Message.Send().ToChatID("oc_3b7ab4655bd6de673cf30cce8ce58b19").SendText(context.Background(), "导入数据出错 截取sku失败,订单号: "+orderDetail.Order.CustomerOrderNo)
				fmt.Println(a, b, c)
			}
			products = append(products, qingliu.Product{
				Name:     skuArr[0],
				Num:      p.Pcs,
				SKU:      sku,
				Price:    p.UnitPrice,
				From:     "import",
				Currency: p.Currency,
			})
		}
		data.Product = products
		result, statusCode := helper.HttpPostJson(data)
		if statusCode != 200 {
			a, b, c := service.LarkCli.Message.Send().ToChatID("oc_3b7ab4655bd6de673cf30cce8ce58b19").SendText(context.Background(), "导入数据出错订单: "+orderDetail.Order.CustomerOrderNo+" 失败")
			fmt.Println(a, b, c)
		}
		fmt.Println(result)

		//执行导入操作
		fmt.Println(orderDetail.ShipperConsignee.ConsigneeFirstName)
	}

	// json.NewEncoder(w).Encode(json.RawMessage(string(rs)))
	json.NewEncoder(w).Encode(json.RawMessage(string(`{"code": 200, "msg": "导入成功"}`)))
}

// 设置4px cookie
func Cookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("cookie")
	service.PXCOOKIE = cookie
	err := helper.PXAlwaysLive(cookie)
	if err != nil {
		json.NewEncoder(w).Encode(json.RawMessage(string(`"code": 500, "msg": "cookie失效"}`)))
	}
	json.NewEncoder(w).Encode(json.RawMessage(string(`{"code": 200, "msg": "cookie更新成功"}`)))
}
