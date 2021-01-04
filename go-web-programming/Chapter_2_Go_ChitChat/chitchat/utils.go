package main

import (
	"errors"
	"fmt"
	"github.com/rwrr/go-web-programming/Chapter_2_Go_ChitChat/chitchat/data"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

func p(a ...interface{}) {
	fmt.Println(a)
}

func init() {
	loadConfig()
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

// 检查用户是否登录并且对应会话存在，否则返回错误
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func version() string {
	return "0.1"
}
