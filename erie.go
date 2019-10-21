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

func ProcessAlert(alert *Alert) {
	alert.Values["from"] = stripSurroundingCarats(alert.Values["from"])

}
