package main

// CPU
// go tool pprof http://localhost:8080/debug/pprof/profile\?seconds\=
// web
// list prepareInfo
// weblist prepareInfo

// Allocation
// go tool pprof http://localhost:8080/debug/pprof/heap
// web prepareInfo

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type CoolInfo struct {
	Message string `json:"message"`
	Sl      []int  `json:"sl"`
}

const duration = time.Second * 5

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/info/", getInfo)
	r.Mount("/debug/", middleware.Profiler())

	go func() {
		someActivity()
	}()

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func getInfo(writer http.ResponseWriter, _ *http.Request) {
	payload := preparePayload()
	_, _ = writer.Write(payload)
}

func preparePayload() []byte {
	info := prepareInfo()
	jsonData, err := json.Marshal(info)
	if err != nil {
		log.Printf("Something bad happened: %s\n", err)
	}
	return jsonData
}

func prepareInfo() CoolInfo {
	sl := make([]int, 0)
	message := "Really cool info!!"
	for i := 0; i < 10000; i++ {
		sl = append(sl, i)
		message += message
	}

	return CoolInfo{
		Message: message,
		Sl:      sl,
	}
}

func someActivity() {
	log.Printf("Some activity start start")
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			action()
		default:
			time.Sleep(duration)
		}
	}
}

func action() {
	for i := 0; i < 10000; i++ {
		_ = prepareInfo()
	}
}
