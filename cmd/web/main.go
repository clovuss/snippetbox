package main

import (
	"flag"
	"log"
	"net/http"
	"os"
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

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":8080", "номер порта на котором запускается веб-сервер")
	flag.Parse()
	logfile, err := os.OpenFile("logfile", os.O_CREATE|os.O_RDWR, 0666)
	defer logfile.Close()
	infoLog := log.New(logfile, "INFO\t", log.LstdFlags|log.Lshortfile)
	errorLog := log.New(logfile, "ERROR\t", log.LstdFlags|log.Lshortfile)
	router := http.NewServeMux()
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	router.Handle("/static/", http.StripPrefix("/static", fileServer))
	infoLog.Printf("Запуск сервера на 127.0.0.1:%s", *addr)

	err = srv.ListenAndServe()

	if err != nil {
		errorLog.Fatal(err)
	}

}
