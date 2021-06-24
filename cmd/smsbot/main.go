package main

import (
	"SmsBot/internal/Database"
	"SmsBot/internal/DiscordBot"
	"fmt"
	"net/http"
	"os"
)

func main() {

	Database.Init()
	go DiscordBot.Init()

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil) //for heroku Port error
}
