package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
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
func request(url string, resultCh chan ResultItem) {
	resp, err := http.Get(url)
	if err != nil {
		resultCh <- ResultItem{Err: err}
		return
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resultCh <- ResultItem{Err: err}
		return
	}

	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		resultCh <- ResultItem{Err: errors.New("Not found")}
		return
	}

	var result string
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
		resultCh <- ResultItem{Err: errors.New("Not found")}
		return
	}

	resultCh <- ResultItem{
		Result: result,
		Lang:   response.Dest,
	}
}

func parseArgs() (string, *Options) {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "dict"
	parser.Usage = "[OPTIONS] SRC"

	args, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
	}

	if opts.From == "" || len(opts.To) == 0 || len(args) != 1 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	return args[0], &opts

}

func main() {

	src, opts := parseArgs()

	toLen := len(opts.To)
	resultCh := make(chan ResultItem, toLen)
	for _, v := range opts.To {
		go request(makeURL(src, opts.From, v), resultCh)
	}

	for i := 0; i < toLen; i++ {
		result := <-resultCh
		if result.Err != nil {
			fmt.Println(result.Err)
			continue
		} else {
			fmt.Printf("%s:%s\n", result.Lang, result.Result)
		}
	}
}

// Options is command line options
type Options struct {
	From string   `short:"f" logn:"from" description:"Source labguage"`
	To   []string `short:"t" long:"to" description:"Destination languages"`
}

// ResultItem is object used to message passing.
type ResultItem struct {
	Lang   string
	Err    error
	Result string
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
