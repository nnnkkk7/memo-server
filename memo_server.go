package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

const saveFile = "memo.text"

func main() {
	print("memo server- [URL] http://localhost:8888/¥n")
	http.HandleFunc("/", readHandler)
	http.HandleFunc("/w", writeHandler)
	http.ListenAndServe(":8888", nil)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadFile(saveFile)
	if err != nil {
		text = []byte("ここにメモを記入してください。")
	}
	htmlText := html.EscapeString(string(text))
	s := "<html>" +
		"<style>textarea { width:90%; height:200px; }</style>" +
		"<form method='POST' action='/w'>" +
		"<textarea name='text'>" + htmlText + "</textarea>" +
		"<input type='submit' value='保存' /></form></html>"
	w.Write([]byte(s))
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["text"]) == 0 {
		w.Write([]byte("投稿してください。"))
		return
	}
	text := r.Form["text"][0]
	ioutil.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save:" + text)
	http.Redirect(w, r, "/", 301)
}
