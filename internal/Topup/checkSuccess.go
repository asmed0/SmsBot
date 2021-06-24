package Topup

import (
	"net/http"
	"smsbot/internal/Database"
	"strconv"
)
var(
	successurl string
	cancelurl string
)
func checkSuccess(w http.ResponseWriter, req *http.Request) {
	successurl = "http://" + req.Host + "/success.html"
	cancelurl = "http://" + req.Host + "/cancel.html"

	qty, ok := req.URL.Query()["qty"]
	sessionID, ok := req.URL.Query()["session_id"]
	discordID, ok := req.URL.Query()["discord_id"]

	if !ok || len(sessionID[0]) < 1 || len(discordID[0]) < 1 {
		http.Redirect(w, req, cancelurl, 301)
	}

	if !Database.ExistsPaymentSecret(sessionID[0]){
		http.Redirect(w, req, cancelurl, 301) //reused payment secret
		return
	}

	Database.AddPaymentSecret(discordID[0], "")
	qty2Add, _ := strconv.Atoi(qty[0])

	go Database.UpdateBalance(qty2Add, discordID[0])
	http.Redirect(w, req, successurl, 301) //successfull topup
}

