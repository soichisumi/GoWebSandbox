package main

import (
	"net/http"
	"sync"
	"html/template"
	"encoding/json"
	"fmt"
	"log"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

// 属性名は大文字でないとexportされない
type result struct {
	Boolean bool 	`json:"boolean"`
	Str string		`json:"str"`
	Num int			`json:"num"`
}

type response struct {
	Res result  `json:"res"`
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//t.once.Do(func(){
	//	t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	//})
	//t.templ.Execute(w, nil)
	obj := response{
		Res: result{
			Boolean: true,
			Num: 114,
			Str: "yoyoyo",
		},
	}
	jsonbytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("")
		return
	}
	fmt.Printf(string(jsonbytes))
	w.Write(jsonbytes)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	if err := http.ListenAndServe(":5005",nil); err != nil {
		log.Fatal("wei")
	}
}

//type Country struct {
//	Name string              `json:"name"`
//	Prefectures []Prefecture `json:"prefectures"`
//}
//
//type Prefecture struct {
//	Name string    `json:"name"`
//	Capital string `json:"capital"`
//	Population int `json:"population"`
//}
//
//func main() {
//	tokyo := Prefecture{Name: "東京都", Capital: "東京", Population: 13482040}
//	saitama := Prefecture{Name: "埼玉県", Capital: "さいたま市", Population: 7249287}
//	kanagawa := Prefecture{Name: "神奈川県", Capital: "横浜市", Population: 9116252}
//	japan := Country{
//		Name:        "日本",
//		Prefectures: []Prefecture{tokyo, saitama, kanagawa},
//	}
//
//	jsonBytes, err := json.Marshal(japan)
//	if err != nil {
//		fmt.Println("JSON Marshal error:", err)
//		return
//	}
//
//	fmt.Println(string(jsonBytes))
//}