package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type ChoTotCar struct {
	url, image, name, price, odo string
}

type LazadaItem struct {
	url, image, name, price string
}

type ShopeeItem struct {
	url, image, name, price string
}

func lazadaScrapper(urlInput string) []LazadaItem {
	// c := colly.NewCollector()
	var lazadaItems []LazadaItem
	var urlPaginate = urlInput
	var i = 1
	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	var nodes []*cdp.Node

	for i < 103 {
		chromedp.Run(ctx,
			chromedp.Navigate(urlPaginate),
			chromedp.Nodes(".Bm3ON", &nodes, chromedp.ByQueryAll),
		)

		// scraping data from each node
		var url, image, name, price string
		for _, node := range nodes {
			chromedp.Run(ctx,
				chromedp.AttributeValue("a", "href", &url, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text(".RfADt", &name, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.AttributeValue("img", "src", &image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				chromedp.Text("span.ooOxS", &price, chromedp.ByQuery, chromedp.FromNode(node)),
			)
			var lazadaItem = LazadaItem{}

			lazadaItem.url = url
			lazadaItem.image = image
			lazadaItem.name = name
			lazadaItem.price = price

			lazadaItems = append(lazadaItems, lazadaItem)
		}
		i++

		urlPaginate = fmt.Sprintf("%s&page=%d", urlInput, i)
		fmt.Printf("Next Loading: %d - %s\n", i, urlPaginate)
	}

	return lazadaItems
}

func choTotScrapper(url string) []ChoTotCar {
	// current iteration
	i := 1
	// max pages to scrape
	limit := 1

	c := colly.NewCollector()
	var choTotCars []ChoTotCar

	// setting a valid User-Agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
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

	c.OnScraped(func(response *colly.Response) {
		// until there is still a page to scrape
		if i < limit {
			pageToScrape := fmt.Sprintf("%s?page=%d", url, i)
			// incrementing the iteration counter
			i++

			fmt.Println("Visiting", pageToScrape)
			// visiting a new page
			c.Visit(pageToScrape)
		}
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)

	return choTotCars
}

func exportLazadaCSV(lazadaItems []LazadaItem) {
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
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, lazadaItem := range lazadaItems {
		// converting a PokemonProduct to an array of strings
		record := []string{
			lazadaItem.url,
			lazadaItem.image,
			lazadaItem.name,
			lazadaItem.price,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}

func exportChototCSV(choTotCars []ChoTotCar) {
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

func exportShopeeCSV(shopeeItems []ShopeeItem) {
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
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, shopeeItem := range shopeeItems {
		// converting a PokemonProduct to an array of strings
		record := []string{
			shopeeItem.url,
			shopeeItem.image,
			shopeeItem.name,
			shopeeItem.price,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	lazadaFlags := flag.String("lazada", "foo", "lazada Flag")
	choTotFlags := flag.String("chotot", "foo", "cho tot Flag")

	flag.Parse()
	var url = ""
	var lazadaItems []LazadaItem
	var chototItems []ChoTotCar

	if *lazadaFlags != "foo" {
		url = *lazadaFlags
		lazadaItems = lazadaScrapper(url)
		exportLazadaCSV(lazadaItems)
	} else {
		if *choTotFlags != "foo" {
			url = *choTotFlags
			chototItems = choTotScrapper(url)
			exportChototCSV(chototItems)
		} else {
			fmt.Println("Unavailable Flag")
		}

	// }
	// url := "https://tiki.vn/dien-thoai-may-tinh-bang/c1789"
	// lazadaItems := tikiScrapper(url)
	// exportLazadaCSV(lazadaItems)

	// url := "https://www.lazada.vn/dien-thoai-di-dong/?spm=a2o4n.home.cate_1.1.d57c3bdcARpMip"
	// lazadaItems := lazadaScrapper(url)
	// exportLazadaCSV(lazadaItems)

	// url := "https://www.nhatot.com/mua-ban-can-ho-chung-cu"
	// choTotCars := choTotScrapper(url)
	// exportChototCSV(choTotCars)
}
