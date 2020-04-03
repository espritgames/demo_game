package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const maxUsers = 10000

func main() {
	levels := make([]int, maxUsers)
	for i := 0; i < maxUsers; i++ {
		levels[i] = rand.Intn(10)
	}

	go func() {
		for {
			level := rand.Intn(maxUsers)
			levels[level+1]++
			time.Sleep(5 * time.Second)
		}
	}()

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/user/{userID}", func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
		if err != nil || userID < 0 {
			w.WriteHeader(400)
			return
		}

		if userID < 1 || userID > maxUsers {
			w.WriteHeader(404)
			return
		}
		level := levels[userID-1]

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": %d, "level": %d}`, userID, level)
	})

	http.ListenAndServe(":8080", r)
}
