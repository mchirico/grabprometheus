package query

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestExtractJson(t *testing.T) {

	dat, err := ioutil.ReadFile("../fixtures/sample.json")
	if err != nil {
		t.Fatalf("%v\n",err)
	}

	p,_ := ExtractJson(dat)

	for _,v := range p.Data.Result {
		if strings.Contains(v.Metric.Hostname,"slice-") != true {
			t.Fatalf("Didn't parse")
		}

	}
	if len(p.Data.Result) != 72 {
		t.Fatalf("Couldn't parse")
	}
	fmt.Printf("%v\n", len(p.Data.Result))
	fmt.Printf("%v\n", p.Data.Result[0].Metric.Job)
	fmt.Printf("%v\n", p.Data.Result[0].Values[0][1])
}