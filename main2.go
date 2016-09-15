package main

import (
	"fmt"
	//"net/url"
	"github.com/parnurzeal/gorequest"
)

type DeviceRegister struct {
	Id 			int		`json:"Id,omitempty"`
	Name 		string 		`json:"name"`
	Appid 		string 		`json:"appid"`
	Uuid 		string 		`json:"uuid"`
	Devicetype 	string 		`json:"devicetype"`
}

func main() {

	//Test
	/*request := gorequest.New()
	_, body, errs := request.Get("http://localhost:8080/v1/app").End()
	if errs != nil {
		fmt.Println(errs)
	}
	fmt.Println(body)*/

	/*queryString := map[string]interface{}{
	  "apikey": "aWSmHBwQoL6IUgEvpBguvsy1FB0HokTj",
	}*/
	device1 := DeviceRegister{Name:"dev-raspberry pi 2", Appid: "7", Uuid: "7e60d58e10386a06ab9cf96d48c8ea06d3ff81e5c1d89cc3f8b0b1cf3ce375", Devicetype: "Raspberry Pi 2"}

	request := gorequest.New()
	_, body, err := request.Post("http://localhost:8080/v1/device").
		Set("ContentType","application/json").
		SendStruct(&device1).
		Param("apikey","aWSmHBwQoL6IUgEvpBguvsy1FB0HokTj").
		End()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)

	/*item := url.Values{}
	item.Set("apikey", "aWSmHBwQoL6IUgEvpBguvsy1FB0HokTj")

	res, err := goreq.Request{
	    Method: "POST",
	    Uri: "http://localhost:8080/v1/device",
	    Accept: "application/json",
    	ContentType: "application/json",
	    QueryString: item,
	    Body: device1,
	}.Do()
	fmt.Println(res.Response, err)*/
}
