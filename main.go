package main

import (
	"net/http"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"log"
	"github.com/dandoh/sdr/auth"
	"github.com/dandoh/sdr/app"
	"github.com/dandoh/sdr/model"
	"fmt"

	"time"
)


const HOUR_TO_TICK int = 11
const MINUTE_TO_TICK int = 38
const SECOND_TO_TICK int = 15

type jobTicker struct {
	t *time.Timer
}

func getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(24 * time.Hour)
	}
	return nextTick.Sub(time.Now())
}

func NewJobTicker() jobTicker {
	return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
	jt.t.Reset(getNextTickDuration())
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.Init()
	model.InitType()

	//Automatically create new report for all user at 2.00 am everyday
	go func(){
		jt := NewJobTicker()
		for {
			<-jt.t.C
				model.CreateTodayReportForAllUsers()
			jt.updateJobTicker()
		}
	}()

	setupServer()

}

func setupMux() *http.ServeMux {
	mux := http.NewServeMux()

	// graphql Handler
	appHandler := app.AppHandler();

	// login Handler
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/signin", auth.LoginFunc)
	mux.HandleFunc("/signup", auth.SignupFunc)

	// add in addContext middlware
	mux.Handle("/graphql", appHandler)

	return mux
}

func setupServer() {
	rootMux := setupMux();
	c := cors.AllowAll().Handler(rootMux);
	fmt.Println(http.ListenAndServe(":8080", c))
}
