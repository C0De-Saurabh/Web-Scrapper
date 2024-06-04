package main

import (
    "fmt"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    url := "https://www.scrapingcourse.com/ecommerce/"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Printf("Status code error: %d %s\n", resp.StatusCode, resp.Status)
        return
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    doc.Find("li.product").Each(func(index int, item *goquery.Selection) {
        
		title := item.Find("h2.woocommerce-loop-product__title").Text()
		imgSrc, exists := item.Find("img").Attr("src")
        if !exists {
            imgSrc = "No image URL found"
        }
        fmt.Printf("Product %d: %s\n", index+1, title)
		fmt.Printf("Image URL: %s\n", imgSrc)
    })
}
