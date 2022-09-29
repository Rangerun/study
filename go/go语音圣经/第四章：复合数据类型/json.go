package main

import (
	"fmt"
	"encoding/json"
)

type ExtData struct {
	HandlerList  Handler `json:"handler"`
	Chn		string `json:"chn"`
}

type Handler struct {
	AreaId	string`json:"area_id"`
	CenterId	string `json:"center_id"`
	SquadId	string `json:"squad"`
}

func main() {
	data := "{\"handler\":{\"area_id\":\"1\",\"center_id\":\"2\",\"group\":\"2\",\"squad\":\"3\"}, \"chn\" : \"aaa\"}"
	
    var extData ExtData
    err := json.Unmarshal([]byte(data), &extData)

    if err !=nil{
        fmt.Println(err)
    }
	
    fmt.Println(extData.HandlerList.AreaId)

}