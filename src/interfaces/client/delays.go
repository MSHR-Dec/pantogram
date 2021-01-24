package client

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Delays struct {
	Names []string
}

func (delays *Delays) Run() {
	rssFeed := delays.getInformation()
	delays.Shape(rssFeed)
}

func (delays *Delays) getInformation() RSSFeed {
	tetsudoCom := "http://api.tetsudo.com/traffic/atom.xml"
	var rssFeed RSSFeed

	req, err := http.NewRequest("GET", tetsudoCom, nil)
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
	return rssFeed
}

func (delays *Delays) Shape(rssFeed RSSFeed) {
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
			delays.Names = append(delays.Names, ts)
		}
	}
}
