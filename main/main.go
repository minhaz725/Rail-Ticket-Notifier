package main

import (
	"Rail-Ticket-Notifier/utils/constants"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

func main() {
	// Create a context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the URL
	fmt.Println("Search Started")
	url := constants.BASE_URL + constants.FROM + "Dhaka" + constants.TO + "Chattogram" + constants.DATE + "17-Dec-2023" + constants.CLASS
	log.Println(url)
	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		log.Fatal(err)
	}

	// Wait for some time (adjust this as needed) to ensure the page has loaded
	// You can use chromedp.Sleep or chromedp.WaitEvent for this purpose
	chromedp.Sleep(5 * time.Second)
	fmt.Println("Search Ended")
	// Extract the page content after it has loaded
	var pageContent string
	if err := chromedp.Run(ctx, chromedp.InnerHTML("html", &pageContent)); err != nil {
		log.Fatal(err)
	}
	// Process and print the page content
	//log.Println(pageContent)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(pageContent))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".single-trip-wrapper").Each(func(i int, element *goquery.Selection) {
		// Extract the train name
		trainName := element.Find(".trip-name h2").Text()
		fmt.Println("Train Name:", trainName)

		// Extract the seat numbers
		element.Find(".seat-available-wrap .all-seats").Each(func(j int, seatElement *goquery.Selection) {
			seatNumber := seatElement.Text()
			fmt.Println("Seat Number:", seatNumber)
		})
	})
}
