package main

import (
	"avitoTest/2_categories/category_item"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"time"
)

/*
Не предствляется способа без привлечения Selenium найти соответствие между
выбранным пунктом меню и ссылкой. Строки, подставляемые при переходе содержатся
в gtm.js, в огромном и уродливом JSON, где все строки лежат в перемешку.
*/

// корневые категории:

var rootPages = []string{
	"https://www.avito.ru/rossiya/transport",
	"https://www.avito.ru/rossiya/nedvizhimost",
	"https://www.avito.ru/rossiya/rabota",
	"https://www.avito.ru/rossiya/predlozheniya_uslug",
	"https://www.avito.ru/rossiya/lichnye_veschi",
	"https://www.avito.ru/rossiya/dlya_doma_i_dachi",
	"https://www.avito.ru/rossiya/bytovaya_elektronika",
	"https://www.avito.ru/rossiya/hobbi_i_otdyh",
	"https://www.avito.ru/rossiya/zhivotnye",
	"https://www.avito.ru/rossiya/dlya_biznesa",
}

func worker(page string, resultResponse chan<- *category_item.CategoryItems, resultError chan<- error) {

	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err := client.Get(page)
	if err != nil {
		resultError <- errors.Wrap(err, "Error on load '"+page+"': ")
		return
	}

	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		resultError <- errors.Wrap(err, "Error on load '"+page+"': ")
		return
	}

	ul := document.Find("div[class^='rubricator-root'] > ul[class^='rubricator-list']").Nodes[0]
	result := recursiveParser(ul)

	resultResponse <- &result
}

func main() {
	resultResponse := make(chan *category_item.CategoryItems, len(rootPages))
	resultError := make(chan error, len(rootPages))

	for _, page := range rootPages {
		go worker(page, resultResponse, resultError)
	}

	var result category_item.CategoryItems
	for i := len(rootPages); i != 0; i-- {
		select {
		case response := <-resultResponse:
			result = append(result, *response...)
		case err := <-resultError:
			//noinspection GoUnhandledErrorResult
			fmt.Fprint(os.Stderr, err)
		}
	}

	if result == nil {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintln(os.Stderr, "Failed to load. Maybe you were blocked on the site?")
		return
	}

	root := category_item.CategoryItem{
		Title:   "Любая категория",
		Href:    "https://www.avito.ru/rossiya",
		Childes: result,
	}

	resultJson, _ := root.MarshalJSON()

	categoriesFile, err := os.Create("categories.json")
	if err != nil {
		//noinspection GoUnhandledErrorResult
		fmt.Fprintln(os.Stderr, "Unable to create ./categories.json : ", err)
		return
	}

	//noinspection GoUnhandledErrorResult,GoNilness
	categoriesFile.Write(resultJson)
}
