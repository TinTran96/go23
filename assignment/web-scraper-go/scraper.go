package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type ChoTotCar struct {
	url, image, name, price, odo string
}

func main() {
	c := colly.NewCollector()
	var choTotCars []ChoTotCar

	// Find and visit all links
	c.OnHTML("li.AdItem_wrapperAdItem__S6qPH", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		var car ChoTotCar

		car.url = e.ChildAttr("a", "href")
		car.price = e.ChildText("p.AdBody_adPriceNormal___OYFU")
		car.odo = e.ChildText("span.AdBody_adItemCondition__ppptn")
		car.image = e.ChildAttr("img.AdThumbnail_thumbnailDefault__mF_DR", "src")
		car.name = e.ChildText("h3")

		choTotCars = append(choTotCars, car)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://xe.chotot.com/mua-ban-oto-toyota-sdcb2")

	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
		"odo",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, choTotCar := range choTotCars {
		// converting a PokemonProduct to an array of strings
		record := []string{
			choTotCar.url,
			choTotCar.image,
			choTotCar.name,
			choTotCar.price,
			choTotCar.odo,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}
