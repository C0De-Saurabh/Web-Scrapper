package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
    baseURL := "https://www.scrapingcourse.com/ecommerce/page/"
    maxPages := 5

    // Open a file for writing
    file, err := os.Create("products.csv")
    if err != nil {
        log.Fatal("Unable to create file:", err)
    }
    defer file.Close()

    // Create a CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write the header
    writer.Write([]string{"Product Title", "Price", "Image URL"})

    // Iterate over the pages
    for i := 1; i <= maxPages; i++ {
        url := baseURL + strconv.Itoa(i)
        fmt.Println("Fetching URL:", url)

        // Create a new HTTP request
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            log.Fatal(err)
        }

        // Set custom headers
        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
        req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
        req.Header.Set("Accept-Language", "en-US,en;q=0.5")

        // Perform the HTTP request
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            log.Fatal(err)
        }
        defer resp.Body.Close()

        // Check if the request was successful
        if resp.StatusCode != 200 {
            log.Fatalf("Failed to fetch the page: %d %s", resp.StatusCode, resp.Status)
        }

        // Parse the HTML content
        doc, err := goquery.NewDocumentFromReader(resp.Body)
        if err != nil {
            log.Fatal(err)
        }

        // Find each product element and extract the data
        doc.Find("li.product").Each(func(index int, item *goquery.Selection) {
            // Extract the product title
            title := item.Find("h2.woocommerce-loop-product__title").Text()
			
			price := item.Find("span.woocommerce-Price-amount").Text()
            if price == "" {
                price = "No price found"
            }

            // Extract the image URL
            imgSrc, exists := item.Find("img").Attr("src")
            if !exists {
                imgSrc = "No image URL found"
            }

            // Write the data to the CSV file
            err := writer.Write([]string{title, price, imgSrc})
            if err != nil {
                log.Fatal("Unable to write record to file:", err)
            }
        })
    }

  
	fmt.Println("Scraping completed and data saved in products.csv file")

}
