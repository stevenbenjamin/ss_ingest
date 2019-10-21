package ss_ingest

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var hVals = make(map[string]int)
var vVals = make(map[string]int)

func TestFileParse(t *testing.T) {
	if false {
		m := make(map[string]int)
		pth := "/Users/steven/work/simple-sense/s3/raw-alerts"
		filepath.Walk(pth, func(path string, f os.FileInfo, err error) error {
			if f.Mode().IsRegular() {
				if bytes, err := ioutil.ReadFile(path); err == nil {
					alert := Parse(string(bytes))
					ProcessAlert(alert)
					//collectValues("text", alert)
					fmt.Printf("\n\n\n\n\nsubject=[%v]\n", alert.Values["subject"])
					fmt.Printf("\nheaderssubject=[%v]\n", alert.Headers["subject"])
					fmt.Printf("\n|%v|\n", alert.Values["text"])
					//for k, _ := range alert.Values {
					//	i, ok := m[k]
					//	if ok {
					// 		m[k] = i + 1
					// 	} else {
					// 		m[k] = 1
					// 	}

					// }

				}
			}

			return nil
		})
		for k, v := range m {
			fmt.Printf("%v = %v\n", k, v)
		}
		printMap("HEADERS", &hVals)
		printMap("VALUES", &vVals)
	}
}

func collectValues(key string, alert *Alert) {
	v := alert.Values[key]
	h := alert.Headers[key]
	addInstanceToMap(v, &vVals)
	addInstanceToMap(h, &hVals)
}
func printMap(label string, mm *map[string]int) {
	m := *mm
	fmt.Println("\n\n--------\n" + label)
	for k, v := range m {
		fmt.Printf("%v\t=\t%v\n", k, v)
	}
}

func addInstanceToMap(key string, m *map[string]int) {
	mf := *m
	i, _ := mf[key]
	mf[key] = i + 1
}

// func TestParseAlert(t *testing.T) {
// 	alert := Parse(string(loadTestFile(t, "EC.COMM-AT-emergycare.org/01:59:50.959Z")))
// 	fmt.Println("\nVALUES")
// 	for k, v := range alert.Values {
// 		fmt.Printf("[%v] = [%v]\n", k, v)
// 	}
// 	fmt.Println("\nHEADERS")
// 	for k, v := range alert.Headers {
// 		fmt.Printf("[%v] = [%v]\n", k, v)
// 	}
// }

// fileList := []string{}
// err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
//     fileList = append(fileList, path)
//     return nil
// })

// for _, file := range fileList {
//     fmt.Println(file)
// }
