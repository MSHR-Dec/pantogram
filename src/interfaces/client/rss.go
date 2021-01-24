package client

type RSSFeed struct {
	Entry []Entry `xml:"entry"`
}

type Entry struct {
	Title string `xml:"title"`
}
