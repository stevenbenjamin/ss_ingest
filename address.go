package ss_ingest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/smartystreets/smartystreets-go-sdk/us-extract-api"
	"github.com/smartystreets/smartystreets-go-sdk/wireup"
	//	"googlemaps.github.io/maps"
	"log"
	"net/http"
	"net/url"
	//	"os"
)

//apiKey: '11fb7c4405d847e0917e92a0f98e5e88',
//provider: 'google',
//	apiKey: 'AIzaSyCX-f8zkkL9thwcuataJOhW9PCikfH-sHo',
//	formatter: 'geoJSON'

var GOOGLE_API_KEY = "AIzaSyCX-f8zkkL9thwcuataJOhW9PCikfH-sHo"
var GOOGLE_API_URL = "https: //maps.googleapis.com/maps/api/geocode/json"

func GoogleVerify(address string) {
	urlstring := fmt.Sprintf("%v?address=%v&key=%v", GOOGLE_API_KEY, url.QueryEscape(address), GOOGLE_API_KEY)
	resp, err := http.Get(urlstring)
	fmt.Printf("Google Response: %v, %v", resp, err)
}

var SMARTY_AUTH_ID = "9ca18209-f049-897b-f53a-93d1b0989258"
var SMARTY_AUTH_TOKEN = "tOZNMbi9p7HhGrdzQ2v7"

func SmartyVerify(address string) {

	log.SetFlags(log.Ltime | log.Llongfile)

	client := wireup.BuildUSExtractAPIClient(
		wireup.SecretKeyCredential("SMARTY_AUTH_ID", "SMARTY_AUTH_TOKEN"),
		//wireup.DebugHTTPOutput(), // uncomment this line to see detailed HTTP request/response information.
	)
	lookup := &extract.Lookup{
		Text:                    address,
		Aggressive:              false,
		AddressesWithLineBreaks: true,
		AddressesPerLine:        1,
	}

	if err := client.SendLookup(lookup); err != nil {
		fmt.Printf("Error sending batch:%v", err)
	} else {
		fmt.Printf("SmartyReturned %v", DumpJSON(lookup))
	}
}

func DumpJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}

	var indent bytes.Buffer

	err = json.Indent(&indent, b, "", "  ")

	if err != nil {
		return err.Error()
	}

	return indent.String()
}

// //Find Address in string
// findAddress: async function(address){
// 	let lookup = new Lookup(address);
// 	lookup.agressive = true;
// 	lookup.addressesHaveLineBreaks = true;
// 	lookupAddressesPerLine = 1;

// 	await smartyClient.send(lookup).then( async function(response) {
// 		var returnAddress = response.result.addresses[0].candidates[0].deliveryLine1
// 		console.log("Smarty Response", returnAddress);
// 		return(returnAddress);
// 	}).catch(console.log);
// },

// verifyAddress: async function(address, returnAddress){
// 	let lookup = new Lookup(address);
// 	lookup.agressive = true;
// 	lookup.addressesHaveLineBreaks = true;
// 	lookupAddressesPerLine = 1;

// 	await smartyClient.send(lookup).then( async function(response) {
// 		returnAddress.a = response.result.addresses[0].candidates[0].deliveryLine1
// 		console.log("Smarty Response", returnAddress);
// 		return(returnAddress.a);
// 	}).catch(console.log);
// },

// //Geolocation Services
// //Google api key
// // AIzaSyB7I2VrZ6tukrlhc_VBYoO2rVXDiZoffPQ
// var geoOptions = {
// 	//provider: 'opencage',
// 	//apiKey: '11fb7c4405d847e0917e92a0f98e5e88',
// 	provider: 'google',
// 	apiKey: 'AIzaSyCX-f8zkkL9thwcuataJOhW9PCikfH-sHo',
// 	formatter: 'geoJSON'
// };

// var geocoder = NodeGeocoder(geoOptions);

// 	geocoder.geocode(address)
// 			.then(function(loc) {
// 				var index = 0;
// 				var bestResult = 0;
// 				var bestIndex = 0;
// 				for(var result in loc) {
// 					// console.log(result);
// 					// console.log(loc[result]);
// 					if(loc[result].extra.confidence > bestResult){
// 						bestResult = loc[result].extra.confidence;
// 						bestIndex = index;
// 					}
// 					index++;
// 				}
// 			var latitude = parseFloat(loc[bestIndex].latitude);
// 				var longitude = parseFloat(loc[bestIndex].longitude);
// 				var range = parseFloat(loc[bestIndex].extra.confidenceKM)

// 				User.find(
// 				{
// 					location: {
// 						$geoIntersects: {
// 							$geometry: {
// 								type: "Point",
// 								coordinates: [longitude, latitude]
// 							}
// 						}
// 					}
