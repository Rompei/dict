package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// URL of api
const URL = "https://glosbe.com/gapi/translate"

func makeURL(src, srcLang, destLang string) string {
	v := url.Values{}
	v.Add("from", srcLang)
	v.Add("dest", destLang)
	v.Add("format", "json")
	v.Add("phrase", src)

	return URL + "?" + v.Encode()
}

// Send request.
func request(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		return
	}

	if lenTuc := len(response.Tuc); lenTuc != 0 {
		res := make([]string, lenTuc)
		count := 0
		for i, v := range response.Tuc {
			if v.Phrase.Text != "" {
				res[i] = v.Phrase.Text
				count = i
			}
		}
		result = strings.Join(res[:count+1], ",")
	}

	if result == "" {
		return "", errors.New("Notfound")
	}
	return
}

func main() {

	// Setting help message.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `dict [OPTIONS] <src>
Options
`)
		flag.PrintDefaults()
	}

	// If this option is exist, translate to Japanese from English.
	from := flag.String("f", "", "From")
	to := flag.String("t", "", "To")
	flag.Parse()
	args := flag.Args()

	if *to == "" || *from == "" || args == nil || len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	result, err := request(makeURL(args[0], *from, *to))
	if err != nil {
		fmt.Println("Not found.")
	} else {
		fmt.Println(result)
	}
}

// Response object.
type Response struct {
	Result string `json:"result"`
	Tuc    []Tuc  `json:"tuc"`
	Phrase string `json:"phrase"`
	From   string `json:"from"`
	Dest   string `json:"dest"`
}

// Tuc object.
type Tuc struct {
	Phrase    Phrase    `json:"phrase"`
	Meanings  []Meaning `json:"meanings"`
	MeaningID int64     `json:"meaningId"`
	Authors   []int64   `json:"authors"`
}

// Phrase object.
type Phrase struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

// Meaning object.
type Meaning struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}
