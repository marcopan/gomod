package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	_ "test.com/mod2/util"
	"text/template"
)

func main() {
	//fmt.Println(util.Util())
	//fmt.Println(util2.Mod2())
	startServer()
}

func startServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello")
	})
	http.HandleFunc("/test", test)

	exec.Command("open", "http://localhost:6066/test").Output()
	http.ListenAndServe(":6066", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("./test.html")
	if err == nil {
		files.Execute(w, "test abc")
	}
	fmt.Println(err)
}
