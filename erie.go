package ss_ingest

import (
	"regexp"
	"strings"
	"time"
)

var dispatchRegex = regexp.MustCompile("32(e|t)(.*)")

var dispatchTypes = map[string]string{"e": " ENGINE", "t": " TOWER"}

func stripSurroundingCarats(s string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, "<"), ">")
}

// Parsed from header TO field: e.g. <eriepa+32e13@alerts.simplesense.io>

// where the address begins with "eriepa+" take the remainder before the "@"
// (here 32e13) and apply the following:

// if the entire string == eci -> "EMERGYCARE"
// 32(e|t)YY -> "ERIE FIRE STATION"
// e -> "ENGINE"
// t -> "TOWER"
// TT -> UNIT #
//
// if there is no match prefix the entire string with "UNIT #"
func parseDispatchUnit(s string) string {
	if s == "eci" {
		return "EMERGYCARE"
	}
	vals := dispatchRegex.FindStringSubmatch(s)
	if len(vals) == 3 {
		dispatchType, _ := dispatchTypes[vals[1]]
		return "ERIE FIRE STATION #" + vals[2] + dispatchType
	}
	return "UNIT # " + s

}

//expect time in the form "Wed, 16 Oct 2019 21:51:37 +2000"
func parseTime(s string) (int64, error) {
	t, err := time.Parse(time.RFC1123Z, s)

	if err == nil {
		return t.Unix(), nil
	}
	return -1, err

}

//parse subject of the form
//SYSTEST:INFOF >TEST12 -  > _:(code) > (type)}
//55B5 >ELECTRICAL HAZ-UNKNOWN SITUA < (code)>(type)

func parseSubject(s string) map[string]string {
	t := strings.Split(s, ">")
	if len(t) == 1 {
		return map[string]string{"type": s}
	}
	t2 := strings.Split(t[0], ":")
	if len(t2) == 1 {
		return map[string]string{"type": t[1], "code": t2[0]}
	}
	return map[string]string{"type": t[1], "code": t2[1]}
}

// parse e.g.

var addrRegex = regexp.MustCompile("(?si)Content-Disposition.*name=\"([^\"]*)\"(.*)")

//return {before, after(trimmed), true if split}
//if there is no split after will be ""
func splitAt(base string, delimiter string) (string, string, bool) {
	out := strings.SplitN(base, delimiter, 2)
	if len(out) == 2 {
		return strings.TrimSpace(out[0]), strings.TrimSpace(out[1]), true
	}
	return base, "", false

}

var cadRegex = regexp.MustCompile("(?s)Cad:\\s?(\\S*)")
var gridsRegex = regexp.MustCompile("(?s)Grids:\\s?(\\S*)")
var mapRegex = regexp.MustCompile("(?s)Map:\\s?(\\S*)")

// ERIE911:53O2  >CITIZEN ASSIST-DOWNED TREE/OBJ
// 1035 E 7TH ST
// XS: EAST AVE
// ERIE CITY
// Map:12102 Grids:,
// Cad: 2019-0000129185

func parseLocation(s string, subject string) *Location {
	loc := &Location{City: "Erie", State: "PA"}
	var base string
	//split off subject
	if _, post, ok := splitAt(base, subject); ok {
		base = post
	} else {
		base = s
	}

	if s1 := cadRegex.FindStringSubmatch(base); len(s1) > 1 {
		loc.Cad = s1[1]
	}
	if s2 := gridsRegex.FindStringSubmatch(base); len(s2) > 1 {
		loc.Grid = s2[1]
	}
	if s3 := mapRegex.FindStringSubmatch(base); len(s3) > 1 {
		loc.Map = s3[1]
		base = strings.SplitN(base, "Map:", 2)[0]
	}
	if preXs, postXs, ok := splitAt(base, "XS:"); ok {
		loc.XS = postXs
		loc.Street = preXs
	}

	return loc
}

func ProcessAlert(alert *Alert) {
	alert.Values["from"] = stripSurroundingCarats(alert.Values["from"])

}
