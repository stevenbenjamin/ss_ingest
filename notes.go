package ss_ingest

import (
	"strings"
)

//Geolocation Services
var geoOptions = map[string]string{
	"provider":  "google",
	"apiKey":    "API_KEY",
	"formatter": "geoJSON"}

var sendGridApiKey = "SENDGRID_API_KEY"

var twilioOptions = map[string]string{
	"accountSid": "...",
	"authToken":  "..."}

// 	var subject = req.body.subject;
// 	var text = req.body.text;
//var from = req.body.from;

//store := if no from, FAIL

var ignoreFrom = map[string]bool{"daniel@nvjumpstarter.com": true, "sys@alerts.simplesense.io": true}

//return store==true if the string is not in the ignoreFrom map
func store(from string) bool {
	_, ok := ignoreFrom[strings.ToLower(from)]
	return !ok
}
