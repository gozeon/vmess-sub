package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func exampleScrape() string {
	// Request the HTML page.
	res, err := http.Get(os.Getenv("VMESS_T_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	var vmessurl, exit = doc.Find(os.Getenv("VMESS_T_SELECT")).Attr(os.Getenv("VMESS_TATTR"))

	if exit {
		return base64.StdEncoding.EncodeToString([]byte(vmessurl))
	} else {
		return ""
	}
}

func main() {
	r := gin.Default()
	fmt.Println(os.Getenv("VMESS_T_URL"))
	fmt.Println(os.Getenv("VMESS_T_SELECT"))
	fmt.Println(os.Getenv("VMESS_TATTR"))
	fmt.Printf("%T", os.Getenv("VMESS_TATTR"))

	r.GET("/vmess", func(c *gin.Context) {

		c.String(http.StatusOK, "%s", exampleScrape())
	})

	r.Run()
}
