package main

import (
	"github.com/leeli73/goFileView/perview"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("I'm Index"))

}


func main(){

	perview.Init("/perview/","no") //初始化

	http.HandleFunc("/index",index)

	http.HandleFunc("/perview/",perview.Handle) //绑定到preview的Handle

	for{
		http.ListenAndServe(":9090", nil)
	}



}