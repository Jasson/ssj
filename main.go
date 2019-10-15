package main

import (
	"fmt"
	"github.com/Jasson/jasonconf/conf"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func init()  {
	conf.Init("./conf/app.ini")

}
func main() {
	file := conf.GetConf("web","STATIC_FILE")
	dir := http.Dir(file)
	//http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(dir)))
	log.Println("file:=", file)
	http.Handle("/", http.StripPrefix("/", http.FileServer(dir)))
	port := conf.GetConf("server", "HTTP_PORT")
	log.Println("【默认项目】服务启动成功 监听端1口 ", port)

	http.ListenAndServe(":"+port, nil)

}
func h2() {
	//dir := http.Dir("../")
	dir := http.Dir("/Users/jason/work/myblog/build/html/")
	http.ListenAndServe(":80", http.FileServer(dir))

}
func hi() {
	//dir := http.Dir("/Users/jason/work/myblog/build/html/")
	dir := http.Dir("../")
	log.Println("dir := ", dir)
	fileHandle := http.FileServer(dir)
	http.Handle("/file/", fileHandle)
	//http.Handle("/log/", http.FileServer(http.Dir("/tmp")))
	http.HandleFunc("/", sayHello) //注册URI路径与相应的处理函数
	port := conf.GetConf("server", "HTTP_PORT")
	log.Println("【默认项目】服务启动成功 监听端1口 ", port)
	er := http.ListenAndServe("0.0.0.0:"+port, nil)
	if er != nil {
		log.Fatal("ListenAndServe: ", er)
	}

}
