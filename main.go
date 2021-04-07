package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/movie42/goScrapper/scrapper"
)

const fileName string = "jobs.csv"

func home(c echo.Context) error {
  return c.File("home.html")
}

func handleScraper(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
  	return c.Attachment(fileName, term + ".csv")
}

func main(){
	e := echo.New()
	e.GET("/", home)
	e.POST("/scrape", handleScraper)
	e.Logger.Fatal(e.Start(":1323"))
}