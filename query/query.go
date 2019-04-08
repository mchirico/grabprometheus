package query

import (
	"encoding/json"
)

type Metric struct {
	Name     string `json:"__name__"`
	Device   string `json:"device"`
	Hostname string `json:"hostname"`
	Instance string `json:"instance"`
	It       string `json:"int"`
	IP       string `json:"ip"`
	Job      string `json:"job"`
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
		return p,err
	}

	return p,nil
}





