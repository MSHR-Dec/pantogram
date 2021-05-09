package delay

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const endpoint = "http://api.tetsudo.com/traffic/atom.xml"

type Delay struct{}

func (d Delay) GetDelayRoutes() (routes []string) {
	delayRoutes := d.requestDelayApi()
	routes = d.extractRouteNames(delayRoutes)
	return
}

func (d Delay) requestDelayApi() (rssFeed rSSFeed) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Print("Error:", err)
	}

	conn := new(http.Client)
	resp, err := conn.Do(req)
	if err != nil {
		log.Print("Error:", err)
	}
	defer resp.Body.Close()

	delayXml, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error:", err)
	}

	xml.Unmarshal(delayXml, &rssFeed)

	return
}

func (d Delay) extractRouteNames(rssFeed rSSFeed) (names []string) {
	for _, v := range rssFeed.Entry {
		rep := regexp.MustCompile(`【.*】`)
		v.Title = rep.ReplaceAllString(v.Title, "")
		titleSlice := strings.Split(v.Title, "・")
		for _, ts := range titleSlice {
			if ts == "特急電車" || ts == "特急列車" {
				continue
			}
			if strings.Contains(ts, "新幹線") {
				continue
			}
			names = append(names, ts)
		}
	}

	return
}
