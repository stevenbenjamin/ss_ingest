package ss_ingest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	contents := loadTestFile(t, "EC.COMM-AT-emergycare.org/01:59:50.959Z")
	alert := Parse(string(contents))
	assert(t, len(alert.Headers) > 0, "headers found")
	assert(t, len(alert.Values) > 0, "values found")
	assert(t, alert.Values["headers"] == "", "no headers field in values")
	// out := Parse(string(contents))
	// if len(out) == 0 {
	// 	t.Error("no fields found")
	// }
	// for k, v := range out {
	// 	fmt.Printf("%v\n\n", k, v)
	// }
	// for k, v := range alert.headers {
	// 	fmt.Printf("%v === %v\n\n", k, v)
	// }
	fmt.Println(reflect.ValueOf(alert.Values).MapKeys())

}
