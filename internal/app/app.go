package app

import (
	_ "app/app/docs"
	"app/internal/config"
	"app/internal/domain/product/storage"
	"app/pkg/client/postgresql"
	"app/pkg/metric"
	"context"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

type App struct {
	config     *config.Config
	router     *httprouter.Router
	httpserver *http.Server
}

func NewApp(config *config.Config) (App, error) {
	log.Print("router init")
	router := httprouter.New()

	log.Print("swagger docs init")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", 301))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)
	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	pgConfig := postgresql.NewPgConfig(config.PostgreSQL.Username, config.PostgreSQL.Password, config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database)
	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		log.Fatal(err)
	}

	productStorage := storage.NewProductStorage(pgClient)
	all, err := productStorage.All(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Print(all)
	return App{
		config: config,
		router: router,
	}, nil
}
func (a *App) Run() {
	a.startHTTP()
}

func (a *App) startHTTP() {
	var listener net.Listener
	if a.config.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		socketPath := path.Join(appDir, "app.sock")
		log.Print("socket path: ", socketPath)
		log.Print("create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("bind application to host: %s and port: %s", a.config.Listen.BindIp, a.config.Listen.Port)
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", a.config.Listen.BindIp, a.config.Listen.Port))
		if err != nil {
			log.Fatal(err)
		}
	}
	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut},
		AllowedOrigins:     []string{"http://localhost:3000", "http://localhost:8000"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Location", "Charset", "Access-Control-Allow-Origin", "Content-type", "content-type"},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{"Location", "Authorization", "Content-Disposition"},
		Debug:              false,
	})
	handler := c.Handler(a.router)
	a.httpserver = &http.Server{
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("application initialized and started")
	if err := a.httpserver.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Print("server shutdown")
		default:
			log.Fatal(err)
		}
	}
	err := a.httpserver.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
