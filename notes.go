package ss_ingest

import (
	"strings"
)

//Geolocation Services
var geoOptions = map[string]string{
	"provider":  "google",
	"apiKey":    "AIzaSyCX-f8zkkL9thwcuataJOhW9PCikfH-sHo",
	"formatter": "geoJSON"}

var sendGridApiKey = "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"

var twilioOptions = map[string]string{
	"accountSid": "AC4d7fc8d3775c2e6b5af084c03fd4acd9",
	"authToken":  "90797b4e54094d0bdfb0a7457766a20c"}

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
