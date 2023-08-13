
# Web Scraper E-commerce

Data Scraping from E-commerce Websites: Chotot, and Tiki.


## Tech Stack
- Go v1.20.6
- [chromedp](https://github.com/chromedp/chromedp)
- [gocolly](https://github.com/gocolly/colly)


## Run

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd my-project
```

Run

For **Lazada** Link
```bash
  go run scraper.go -lazada https://www.lazada.vn/dien-thoai-di-dong/?spm=a2o4n.home.cate_1.1.d57c3bdcARpMip
```

For **ChoToT** Link
```bash
  go run scraper.go -chotot https://www.chotot.com/mua-ban-laptop-tp-ho-chi-minh
```


## Instruction

Currently, our tool support Lazada and Chotot site by search URL or category URL
For example:

https://www.lazada.vn/catalog/?q=laptop&_keyori=ss&from=input&spm=..search.go.

https://www.lazada.vn/trang-phuc-nam/?spm=a2o4n.home.cate_9.1.20183bdcbcdT9p

https://xe.chotot.com/mua-ban-xe-may-honda-tp-ho-chi-minh-sdmb1

https://www.chotot.com/mua-ban-dien-thoai-tp-ho-chi-minh

...

Information on Chotot is easy to scrape, except the Thumbnail Image is lazy loading (can only scrape first 4 item thumbnails).

Scraping products from Lazada with Colly can be difficult due to its dynamic content loading. Lazada relies on JavaScript for page rendering or uses it to perform API calls and retrieve data asynchronously

The solution is chromedp, which is a library that provides browser capabilities and allows you to load a web page in a special browser with no GUI. You can then instruct the headless browser to mimic user interactions.

Scraping Lazada with Chromedp takes longer than using GoColly.
Loading this [page](https://www.lazada.vn/dien-thoai-di-dong/?spm=a2o4n.home.cate_1.1.d57c3bdcARpMip) of Lazada with 102 pages with 40 record per page, take 41'34.4178517s.
## Authors

- [@TinTran96](https://github.com/TinTran96)

