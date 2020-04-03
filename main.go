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
		levels[i] = 1 + rand.Intn(9)
	}

	go func() {
		for {
			userID := rand.Intn(maxUsers)
			levels[userID+1]++
			time.Sleep(100 * time.Millisecond)
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

		// Simulate response delay
		time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(900)))
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"id": %d, "level": %d}`, userID, level)
	})

	http.ListenAndServe(":8080", r)
}
