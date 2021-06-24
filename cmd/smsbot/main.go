package main

import (
	"fmt"
	"net/http"
	"os"
	"smsbot/internal/Database"
	"smsbot/internal/DiscordBot"
)

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil) //for heroku Port error

	Database.Init()
	DiscordBot.Init()
}
