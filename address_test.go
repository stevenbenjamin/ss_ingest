package ss_ingest

import (
	"testing"
	"time"
)

func TestSmartyParseAddress(t *testing.T) {
	defer timer(time.Now(), "smarty")
	SmartyVerify("1035 E 7TH ST City, Erie, PA")
}

func TestGoogleParseAddress(t *testing.T) {
	defer timer(time.Now(), "google")
	GoogleVerify("1035 E 7TH ST City, Erie, PA")
}
