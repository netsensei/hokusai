package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Options struct {
	Host string
	Port int
}

func Start(opts Options) error {
	router := chi.NewRouter()

	router.Mount("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	addr := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	output := fmt.Sprintf("The hokusai server is running at http://%s.", addr)
	fmt.Println(output)
	return http.ListenAndServe(addr, router)
}
