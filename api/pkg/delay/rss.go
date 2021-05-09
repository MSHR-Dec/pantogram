package delay

type rSSFeed struct {
	Entry []Entry `xml:"entry"`
}

type Entry struct {
	Title string `xml:"title"`
}
