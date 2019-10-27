package ss_ingest

import (
	"fmt"
	"testing"
	"time"
)

func TestByteParse(t *testing.T) {
	str := loadTestFile(t, "google_address_response.json")
	bytes := []byte(str)
	loc, err := parseGoogleLocation(bytes)
	fmt.Printf("\n================\n%+v\n================\nerr=%v\n", loc, err)

}

func TestGoogleParseAddress(t *testing.T) {
	if false {
		defer timer(time.Now(), "google")
		GoogleVerify("1035 E. 7TH ST City, Erie, PA")
	}
}
