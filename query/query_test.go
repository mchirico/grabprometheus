package query

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestExtractJson(t *testing.T) {

	dat, err := ioutil.ReadFile("../fixtures/sample.json")
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	p, _ := ExtractJson(dat)

	for _, v := range p.Data.Result {
		if strings.Contains(v.Metric.Hostname, "slice-") != true {
			t.Fatalf("Didn't parse")
		}

	}
	if len(p.Data.Result) != 72 {
		t.Fatalf("Couldn't parse")
	}
	fmt.Printf("%v\n", len(p.Data.Result))
	fmt.Printf("%v\n", p.Data.Result[0].Metric.Job)
	//fmt.Printf("%v\n", p.Data.Result[0].Metric.Hostname)
	//fmt.Printf("%v\n", p.Data.Result[1].Metric.Hostname)
	fmt.Printf("%v\n", p.Data.Result[0].Values[0][1])
	fmt.Printf("%v\n", p.Data.Result[0].Values[0][0])
	fmt.Printf("%s\n", GetTime(p.Data.Result[0].Values[1][0]))
	fmt.Printf("len: %d\n", len(p.Data.Result[0].Values))
}

func TestWriteCSV(t *testing.T) {

	dat, err := ioutil.ReadFile("../fixtures/5d")
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	p, _ := ExtractJson(dat)

	WriteCSV("junk.csv", p)

}

func sTestPullData(t *testing.T) {
	url := `http://vovprm-po-1p.sys.comcast.net:9090/api/v1/query?query=cos_device_network_receive_drop{device=%22slicestor%22,%20int=%22bond0%22}[1h]`
	url = `http://vovprm-po-1p.sys.comcast.net:9090/api/v1/query?query=cos_device_cpu{device=%22slicestor%22}[1h]`
	data, _ := PullData(url)
	p, _ := ExtractJson(data)
	fmt.Printf("%v\n", p.Data.Result[0].Metric.Job)

	for _, v := range p.Data.Result {
		if strings.Contains(v.Metric.Hostname, "slice-") != true {
			t.Fatalf("Didn't parse")
		}

		mtime := int64(v.Values[0][0].(float64)*1000) * 1000000
		value := v.Values[0][1]
		host := v.Metric.Hostname
		fmt.Printf("%s: %s, %v\n", host, time.Unix(0, mtime), value)

	}

}

func sTest2(t *testing.T) {
	url := `http://vovprm-po-1p.sys.comcast.net:9090/api/v1/query_range?query=cos_device_network_receive_drop{device=%22slicestor%22,%20int=%22bond0%22}&start=2019-04-01T20:10:30.781Z&end=2019-04-06T20:11:00.781Z&step=20m`

	//url = `http://vovprm-po-1p.sys.comcast.net:9090/api/v1/query?query=cos_device_cpu{device=%22slicestor%22}[1h]`
	data, _ := PullData(url)
	p, _ := ExtractJson(data)
	fmt.Printf("%v\n", p.Data.Result[0].Metric.Job)

	for _, v := range p.Data.Result {
		if strings.Contains(v.Metric.Hostname, "slice-") != true {
			t.Fatalf("Didn't parse")
		}

		if v.Metric.Hostname == "slice-ch2-a10p.sys.comcast.net" {
			fmt.Printf("Value: %v\n", v.Values[0][1])
		}
	}

}
