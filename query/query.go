package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Metric struct {
	Name     string `json:"__name__"`
	Device   string `json:"device"`
	Hostname string `json:"hostname"`
	Instance string `json:"instance"`
	It       string `json:"int"`
	IP       string `json:"ip"`
	Job      string `json:"job"`
	Mode     string `json:"mode"`
}

type R struct {
	Metric Metric          `json:"metric"`
	Values [][]interface{} `json:"values"`
}

type D struct {
	ResultType string `json:"resultType"`
	Result     []R    `json:"result"`
}

type prom struct {
	Status string `json:"status"`
	Data   D      `json:"data"`
}

type Member struct {
	Name           string
	Age            int
	SecretIdentity string
	Powers         []string
	Stuff          string
}

func ExtractJson(input []byte) (prom, error) {
	p := prom{}
	err := json.Unmarshal(input, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func PullData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err

}

func GetTime(i interface{}) time.Time {

	mtime := int64(i.(float64)*1000) * 1000000
	return time.Unix(0, mtime)

}

func WriteCSV(file string, p prom) error {

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Find max
	max := 0
	for _, v := range p.Data.Result {
		if max < len(v.Values) {
			max = len(v.Values)
		}
	}

	for idx := 0; idx < max; idx++ {
		if idx == 0 {
			for i, v := range p.Data.Result {
				if i == 0 {
					f.WriteString("date")
				}
				f.WriteString(fmt.Sprintf(",%s", v.Metric.Hostname))
			}
			f.WriteString("\n")
		}

		for i, v := range p.Data.Result {

			if i == 0 {
				if len(v.Values) > idx {
					dateTime := fmt.Sprintf("%s", GetTime(v.Values[idx][0]))
					f.WriteString(dateTime)
				}
			}

			if len(v.Values) > idx {
				if len(v.Values[idx]) >= 2 {
					f.WriteString(fmt.Sprintf(",%v", v.Values[idx][1]))
				}

			} else {
				f.WriteString(",")
			}

		}
		f.WriteString("\n")
	}

	return nil

}
