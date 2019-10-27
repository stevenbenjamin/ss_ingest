package ss_ingest

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
)

var GOOGLE_API_KEY = "AIzaSyCX-f8zkkL9thwcuataJOhW9PCikfH-sHo"
var GOOGLE_API_URL = "https://maps.googleapis.com/maps/api/geocode/json"

//expected response
// {
//    "results" : [
//       {
//          "address_components" : [
//             {
//                "long_name" : "1035",
//                "short_name" : "1035",
//                "types" : [ "street_number" ]
//             },
//             {
//                "long_name" : "East 7th Street",
//                "short_name" : "E 7th St",
//                "types" : [ "route" ]
//             },
//             {
//                "long_name" : "Zone 1",
//                "short_name" : "Zone 1",
//                "types" : [ "neighborhood", "political" ]
//             },
//             {
//                "long_name" : "Erie",
//                "short_name" : "Erie",
//                "types" : [ "locality", "political" ]
//             },
//             {
//                "long_name" : "Erie County",
//                "short_name" : "Erie County",
//                "types" : [ "administrative_area_level_2", "political" ]
//             },
//             {
//                "long_name" : "Pennsylvania",
//                "short_name" : "PA",
//                "types" : [ "administrative_area_level_1", "political" ]
//             },
//             {
//                "long_name" : "United States",
//                "short_name" : "US",
//                "types" : [ "country", "political" ]
//             },
//             {colo
//                "long_name" : "16503",
//                "short_name" : "16503",
//                "types" : [ "postal_code" ]
//             },
//             {
//                "long_name" : "1511",
//                "short_name" : "1511",
//                "types" : [ "postal_code_suffix" ]
//             }
//          ],
//          "formatted_address" : "1035 E 7th St, Erie, PA 16503, USA",
//          "geometry" : {
//             "bounds" : {
//                "northeast" : {
//                   "lat" : 42.13705849999999,
//                   "lng" : -80.06083919999999
//                },
//                "southwest" : {
//                   "lat" : 42.1369383,
//                   "lng" : -80.0609764
//                }
//             },
//             "location" : {
//                "lat" : 42.1369958,
//                "lng" : -80.060903
//             },
//             "location_type" : "ROOFTOP",
//             "viewport" : {
//                "northeast" : {
//                   "lat" : 42.1383473802915,
//                   "lng" : -80.05955881970849
//                },
//                "southwest" : {
//                   "lat" : 42.1356494197085,
//                   "lng" : -80.06225678029149
//                }
//             }
//          },
//          "place_id" : "ChIJYUi-psl_LYgRmQwdoqhtJLI",
//          "types" : [ "premise" ]
//       }
//    ],
//    "status" : "OK"
// }

func parseGoogleLocation(raw []byte) (*Location, error) {
	loc := new(Location)
	// "OK" indicates that no errors occurred; the address was successfully parsed and at least one geocode was returned.
	// "ZERO_RESULTS" indicates that the geocode was successful but returned no results. This may occur if the geocoder was passed a non-existent address.
	// OVER_DAILY_LIMIT billing error
	// "OVER_QUERY_LIMIT" indicates that you are over your quota.
	// "REQUEST_DENIED" indicates that your request was denied.
	// "INVALID_REQUEST" generally indicates that the query (address, components or latlng) is missing.
	// "UNKNOWN_ERROR"
	status := gjson.GetBytes(raw, "status").Str
	if status != "OK" {
		return loc, errors.New(fmt.Sprintf("Google geocoding error:%v", status))
	}

	//friends.#(nets.#(=="fb"))#.first

	vars := gjson.GetManyBytes(raw,
		"results.0.geometry.location.lat",
		"results.0.geometry.location.lng",
		"results.0.formatted_address",
		`results.0.address_components.#(types.#(=="street_number")).long_name`,
		`results.0.address_components.#(types.#(=="route")).long_name`,
		`results.0.address_components.#(types.#(=="locality")).long_name`,
		`results.0.address_components.#(types.#(=="administrative_area_level_1")).long_name`,
		`results.0.address_components.#(types.#(=="postal_code")).long_name`,
	)

	loc.Latitude = vars[0].Num
	loc.Longitude = vars[1].Num
	loc.FormattedAddress = vars[2].Str
	loc.Street = fmt.Sprintf("%v %v", vars[3].Str, vars[4].Str)
	loc.City = vars[5].Str  //city.Str
	loc.State = vars[6].Str //state.Str
	loc.Zip = vars[7].Str   //zip.Str
	return loc, nil
}

func GoogleVerify(address string) (*Location, error) {
	urlstring := fmt.Sprintf("%v?address=%v&key=%v", GOOGLE_API_URL, url.QueryEscape(address), GOOGLE_API_KEY)
	if resp, err := http.Get(urlstring); err == nil {
		if bytes, e2 := parseResponse(resp); e2 == nil {
			return parseGoogleLocation(bytes)

		} else {
			return nil, e2
		}
	}
	return nil, nil
}

func parseResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
