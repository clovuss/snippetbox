package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path !="/" {
		http.NotFound(w,r)

		return
	}
	_, err := w.Write([]byte("Привет из сниппетбокса"))
	if err != nil {
		fmt.Println(err)
	}

}

func showSnippet (w http.ResponseWriter, r *http.Request)  {
	_, err := w.Write([]byte("Показали сниппет"))
	if err != nil {
		fmt.Println(err)
	}
	
}
func createSnippet (w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST"{
		w.Header().Set("allis", http.MethodPost)

		http.Error(w,"Запрещtно", 405 )
		return
	}
	w.Write([]byte("Создаем заметку!\n"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/snippet", showSnippet)
	router.HandleFunc("/snippet/create", createSnippet)
log.Println("Запуск сервера на 127.0.0.1:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}

}
