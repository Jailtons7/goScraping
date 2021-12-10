package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("#phrasesList", func(e *colly.HTMLElement) {
		writer.Write([]string{"Frase", "Autor"})
		e.ForEach(".thought-card", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("p.frase"),
				el.ChildText("span.autor"),
			})
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://www.pensador.com/programacao/")
}
