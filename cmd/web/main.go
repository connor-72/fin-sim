package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/form/v4"
)

type application struct {
	formDecoder *form.Decoder
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	formDecoder := form.NewDecoder()

	app := &application{
		formDecoder: formDecoder,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	logger.Info("starting server", "addr", srv.Addr)
}
