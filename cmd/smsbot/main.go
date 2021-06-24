package main

import (
	"fmt"
	"net/http"
	"os"
	"smsbot/internal/Database"
	"smsbot/internal/DiscordBot"
	"smsbot/internal/Topup"
)

func main() {

	Database.Init()
	go DiscordBot.Init()

	go Topup.Init()

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil) //for heroku Port error
}
