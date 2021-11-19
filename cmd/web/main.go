package main

import (
	"flag"
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
	addr := flag.String("addr", ":8080", "номер порта на котором запускается веб-сервер")
	flag.Parse()
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/snippet", showSnippet)
	router.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	router.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Printf("Запуск сервера на 127.0.0.1:%s", *addr)

	err := http.ListenAndServe(*addr, router)

	if err != nil {
		fmt.Println(err)
	}

}
