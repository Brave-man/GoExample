package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const provinceApi = "https://api.jisuapi.com/area/province?appkey=86e6d3c18c5b0d65"

type Province struct {
	ID         int    `json:"id"`         // ID
	Name       string `json:"name"`       // 名称
	ParentID   int    `json:"parentid"`   // 上级ID
	ParentName string `json:"parentname"` // 上级名称
	AreaCode   string `json:"areacode"`   // 区号
	ZipCode    string `json:"zipcode"`    // 邮编
	Depth      int    `json:"depth"`      // 区域等级(深度) 冗余字段，用来查找
}

type Result struct {
	Status int        `json:"status"`
	Msg    string     `json:"msg"`
	Result []Province `json:"result"`
}

func main() {
	resp, err := http.Get(provinceApi)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var r Result
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(r)
}
