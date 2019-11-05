package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//func getVivoReqData(click entitys.Click) []byte {
//	type DataList struct {
//		UserIDType string `json:"userIdType"`
//		UserID     string `json:"userId"`
//		CvType     string `json:"cvType"`
//		CvTime     int    `json:"cvTime"`
//		CreativeId string `json:"creativeId"`
//	}
//
//	type RequestStr struct {
//		SrcType  string `json:"srcType"`
//		PkgName  string `json:"pkgName"`
//		SrcID    string `json:"srcId"`
//		DataList DataList `json:"dataList"`
//	}
//
//	var reqStr RequestStr
//	reqStr.SrcType = "APP"
//	reqStr.PkgName = "me.yidui"
//	//bjmlkj数据源ID： ds-201910123719    vivonew2
//	//yidui01数据源ID：ds-201910123319    vivonew1
//	if click.MarketType == "vivonew1" {
//		reqStr.SrcID = "ds-201910123319"
//	} else if click.MarketType == "vivonew2" {
//		reqStr.SrcID = "ds-201910123719"
//	}
//	reqStr.DataList.UserIDType = "IMEI_MD5"
//	reqStr.DataList.UserID = strings.ToLower(click.DeviceId)    //imei_md5
//	reqStr.DataList.CvType = "REGISTER"
//	reqStr.DataList.CvTime = int(click.ClickAt.UnixNano()/1000000)    //ts
//	reqStr.DataList.CreativeId = click.Cid
//
//	data, err := json.MarshalIndent(reqStr, "", "    ")
//	if err != nil {
//		log.Println("vivo callback JSON marshling failed:", err)
//	}
//	// fmt.Printf("%s\n", data)
//	return data
//}

//func getVivoSign(click entitys.Click) string {
//	data := getVivoReqData(click)
//	var apiKey string
//		apiKey = "oof+KSv0gOk="
//	req := fmt.Sprintf("apiKey=%s&requestStr=%s", apiKey, data)
//	//log.Println("vivo getVivoSign request field: ", req)
//	url := "https://ad-market.vivo.com.cn/test/sign"
//	resp, err := http.Post( url,
//		"application/x-www-form-urlencoded",
//		strings.NewReader(req))
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//	//log.Println("get_vivo_sign_response:", string(body))
//	return string(body)
//}
//
//func vivoUpload(click entitys.Click) {
//	data := getVivoReqData(click)
//	sign := getVivoSign(click)
//	var uuid string
//	if click.MarketType == "vivonew1" {
//		uuid = "ab2547084ebf1ec187094532c83ec05c"
//	} else if click.MarketType == "vivonew2" {
//		uuid = "V/6STaespfe4nXGbKtgwa6sQfljJb17Y"
//	}
//	req := fmt.Sprintf("apiUuid=%s&requestStr=%s&sign=%s", uuid, data, sign)
//	//log.Println("vivo upload field: ", req)
//	url := "https://ad-market.vivo.com.cn/v1/advertiser/behavior/post"
//	resp, err := http.Post( url,
//		"application/x-www-form-urlencoded",
//		strings.NewReader(req))
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println("vivo ioutil read err: ", err)
//	}
//	log.Println("vivo_callback_response:", string(body))
//}
func main() {

	url := "https://ad-market.vivo.com.cn/test/sign"

	payload := strings.NewReader("requestStr=%7B%22startDate%22%3A%2220190716%22%2C%22endDate%22%3A%2220190720%22%2C%22page%22%3A1%2C%22pageSize%22%3A40%7D&apiKey=oof%2BKSv0g0k%3D")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "d73f16ad-bde0-4749-881c-75c2e98882f7")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
