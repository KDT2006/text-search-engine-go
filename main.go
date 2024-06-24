package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()
	fmt.Println("Full text search is in progress...")
}
