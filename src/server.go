package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	RequestId string
	Result string
}

type Result struct {
	Response Response
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//接收请求参数
	requestId := r.Header.Get("X-TC-RequestId")
	log.Println("requestId:", requestId)
	name := r.FormValue("name")
	log.Println("name:", name)
	//返回响应
	result := new(Result)
	result.Response.RequestId = requestId
	res,_ := json.Marshal(result)
	w.Write(res)
}

func main () {
	fmt.Printf("listen:%v","13000...")
	http.HandleFunc("/mock", HelloHandler)
	http.ListenAndServe(":13000", nil)
}