package main

import (
	"net/http"
	. "core"
	_ "mvc/controller"
	_ "data-base"
	"log"
	"fmt"
)

type WebServer struct{}

func (ws WebServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	Router.Execute(res , req )
}

func main() {
	fmt.Println("项目启动：8888")
	err := http.ListenAndServe(":8888", WebServer{})
	if err != nil {
		log.Fatal(err)
	}
}
