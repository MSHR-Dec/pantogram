package client_test

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	. "github.com/MSHR-Dec/pantogram/src/interfaces/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delays", func() {
	Context("Delays.shape()", func() {
		It("returns delaying routes except the specific routes", func() {
			delays := &Delays{}
			rssFeed := createRSSFeed("testdata/rss.xml")
			delays.Shape(rssFeed)
			Expect(delays).To(Equal(&Delays{Names: []string{
				"高山線",
				"太多線",
				"はるか",
				"くろしお",
				"木次線",
				"芸備線",
				"山陽本線",
				"特急ラピート",
				"特急サザン",
				"島原鉄道線",
				"大井川本線",
				"SL急行",
			}}))
		})
	})
})

func createRSSFeed(src string) RSSFeed {
	var rssFeed RSSFeed
	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		log.Fatal(err)
	}
	return rssFeed
}
