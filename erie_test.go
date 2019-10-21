package ss_ingest

import (
	"fmt"
	"testing"
)

func TestParseDispatchUnit(t *testing.T) {
	assert(t, parseDispatchUnit("eci") == "EMERGYCARE", "parse eci")
	assert(t, parseDispatchUnit("32exyz") == "ERIE FIRE STATION #xyz ENGINE", "parse32eXXX")
	assert(t, parseDispatchUnit("32txyz") == "ERIE FIRE STATION #xyz TOWER", "parase32tXX")
	assert(t, parseDispatchUnit("32zxyz") == "UNIT # 32zxyz", "parse32???")
	assert(t, parseDispatchUnit("nomatch") == "UNIT # nomatch", "parse no match")

}

func TestProcessFrom(t *testing.T) {
	a := Alert{Values: map[string]string{"from": "<xyz>"}}
	ProcessAlert(&a)
	assert(t, a.Values["from"] == "xyz", "parse clean <from>")
	ProcessAlert(&a)
	assert(t, a.Values["from"] == "xyz", "from without <> is unaltered")
}

func TestParseTime(t *testing.T) {
	t1, _ := parseTime("Wed, 16 Oct 2019 21:51:37 +2000")
	assert(t, t1 == int64(1571190697), "time parsed")
}

func TestParseSubject(t *testing.T) {
	m := parseSubject("PRE>POST")
	assert(t, m["type"] == "POST" && m["code"] == "PRE", "OK")
	m = parseSubject("XYZ:PRE>POST")
	assert(t, m["type"] == "POST" && m["code"] == "PRE", "OK")
	m = parseSubject("POST")
	assert(t, m["type"] == "POST" && m["code"] == "", "OK")
}

func TestParseLocation(t *testing.T) {

	var loc = parseLocation(`
 ERIE911:53O2  >CITIZEN ASSIST-DOWNED TREE/OBJ
 1035 E 7TH ST
 XS: EAST AVE
 ERIE CITY
 Map: 12102 Grids:,
 Cad: 2019-0000129185`, "ERIE911:53O2  >CITIZEN ASSIST-DOWNED TREE/OBJ")

	fmt.Printf("%+v", loc)
}
