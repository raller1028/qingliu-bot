package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"qingliuBot/model/qingliu"
)

func HttpPostJson(data qingliu.QSource) (string, int) {

	jsonStr, _ := json.Marshal(&data)

	// url := "https://qingflow.com/api/qsource/b65d0563-c213-45e5-8318-d0c9d9028ad6" //测试
	url := "https://qingflow.com/api/qsource/1267b993-da3a-4ba2-9b64-5808eca3fdf2" //正式
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return "", 0
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), statuscode
}
