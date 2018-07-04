package helpers

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"

	"github.com/a-fis/insta-stats/app/models"
)

var regexForSharedData *regexp.Regexp

func init() {
	regexForSharedData, _ = regexp.Compile(`<script type="text/javascript">window._sharedData = (.*?);</script>`)
}

// GetHTML ...
func GetHTML(urlstr, ua string) (string, error) {

	if ua == "" {
		ua = "Googlebot"
	}
	req, _ := http.NewRequest("GET", urlstr, nil)
	// Set User-Agent to Googlebot
	req.Header.Set("User-Agent", ua)

	//gzip
	req.Header.Add("Accept-Encoding", "gzip")
	//req.Header.Add("Accept-Encoding", "gzip, deflate")

	// dump req header
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)
	// New Client
	tr := &http.Transport{
		//MaxIdleConns:       10,
		//IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	cl := &http.Client{
		Transport: tr,
	}
	// Send request
	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.Header.Get("Content-Encoding") == "gzip" {

		zipread, e := gzip.NewReader(resp.Body)
		if e != nil {
			log.Print(e)
			return "", e
		}
		defer zipread.Close()

		reader := bufio.NewReader(zipread)
		var part []byte
		ret := ""

		for {
			if part, _, err = reader.ReadLine(); err != nil {
				break
			}
			ret += string(part)
		}
		log.Printf("unzipped resp body size = %d", len(ret))
		return ret, nil
	}

	//log.Print("b, err := ioutil.ReadAll(resp.Body)")
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Printf("resp body size = %d", len(string(b)))
	return string(b), nil

}

// GetIGJsonData ...
func GetIGJsonData(urlstr, ua string) (models.IGJsonData, int64, error) {
	jsonData := models.IGJsonData{}

	log.Print("helpers GetIGJsonData")

	body, err := GetHTML(urlstr, ua)

	// Get HTML
	if err != nil {
		log.Print(err)
		return jsonData, 503, err
	}
	// Find JSON
	//regexForSharedData, _ = regexp.Compile(`<script type="text/javascript">window._sharedData = (.*?);</script>`)
	v := regexForSharedData.FindSubmatch([]byte(body))
	if len(v) <= 1 {
		//
		//TODO: error handling more...
		log.Print("")
		return jsonData, 503, fmt.Errorf("FindSubmatch Error")

	}
	jsonStr := string(v[1])

	//parse JSON
	err = json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		//log.Print(jsonStr)
		log.Print(err)
		log.Print(jsonStr)
		return jsonData, 503, err
	}

	log.Print("Find JSON: OK")
	return jsonData, 200, nil
}
