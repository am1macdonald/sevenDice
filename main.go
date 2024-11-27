package main

import (
	"context"
	"encoding/json"
	"github.com/am1macdonald/sevenDice/internal/apiConfig"
	"github.com/am1macdonald/sevenDice/internal/database"
	"github.com/am1macdonald/sevenDice/internal/mailer"
	"github.com/am1macdonald/sevenDice/internal/router"
	"github.com/am1macdonald/sevenDice/internal/templates"
	"github.com/jackc/pgx/v5"
	"github.com/valkey-io/valkey-go"
	"log"
	"net/http"
	"os"
	"time"
)

var port string = os.Getenv("HOST_PORT")

func jsonResponse(payload any, code int, w http.ResponseWriter) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		_, err = w.Write([]byte("internal error"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(bytes)
	if err != nil {
		log.Println(err)
	}
	return
}

func errorResponse(err error, code int, w http.ResponseWriter) {
	log.Printf("%v\n\n", err)
	jsonResponse(err.Error(), code, w)
}

func connectDB(ctx context.Context) (*pgx.Conn, *database.Queries) {
	conn, err := pgx.Connect(
		ctx,
		"postgres://sevendice:sevendice@db:5432/sevendice",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	return conn, database.New(conn)
}

func connectCache() *valkey.Client {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{"valkey:6379"},
	})

	if err != nil {
		panic(err)
	}

	log.Println("Connected to cache")
	return &client
}

func init() {}

func main() {
	ctx := context.Background()
	conn, Queries := connectDB(ctx)
	defer conn.Close(ctx)

	client := connectCache()
	m := mailer.New(os.Getenv("SMTP_KEY"), os.Getenv("SMTP_EMAIL"))

	t, err := templates.NewTemplateHandler()
	if err != nil {
		log.Fatal(err)
	}

	appRouter := router.NewRouter(client, Queries, m, t)

	cfg := apiConfig.New(appRouter)

	mux := http.NewServeMux()
	mux.HandleFunc("/", cfg.Router.HandleRoute)

	// Initialize server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server starting on http://localhost%s", server.Addr)

	log.Fatal(server.ListenAndServe())
}
