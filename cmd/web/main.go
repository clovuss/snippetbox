package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	file, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := file.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := file.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return file, nil
}
func main() {

	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/snippet", showSnippet)
	router.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
	router.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Запуск сервера на 127.0.0.1:8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err)
	}

}
