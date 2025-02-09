package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/OlegSchwann/core-services-train/3_categories_moscow/category_generator"
	"github.com/OlegSchwann/core-services-train/3_categories_moscow/category_item"
	"github.com/OlegSchwann/core-services-train/3_categories_moscow/statistics_format"
	"github.com/OlegSchwann/core-services-train/3_categories_moscow/worker_creator"
)

// Узнать количество товаров категории productCategory в месте location
// можно перейдя на страницу:
// Get https://www.avito.ru/{location}{productCategory}
// location = "sankt-peterburg"
// productCategory = "/mototsikly_i_mototehnika/mototsikly
// и найти число:
// let numberOfGoods = +document.querySelector('.page-title-count').innerText.replace(/\s+/g, '')

var (
	path     string
	location string
)

func init() {
	flag.StringVar(&path, "categories_path", "./categories.json", "file with tree of categories, like \"https://github.com/OlegSchwann/core-services-train/blob/6327c62465ba12d74fb2b87a3ad5812da1eb3f9d/2_categories/categories.json\".")
	flag.StringVar(&location, "location", "moskva", "Location in avito url, \"moskva\", for example: \"https://www.avito.ru/{moskva}/mototsikly_i_mototehnika/mototsikly\"")
	flag.Parse()
}

func main() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to open file "+path+":", err)
	}

	var category category_item.CategoryItem
	categoryJson, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("unable to extract ")
	}

	//noinspection GoUnhandledErrorResult,GoNilness
	file.Close()

	err = category.UnmarshalJSON(categoryJson)
	if err != nil {
		log.Fatal()
	}

	results, errs := worker_creator.WorkerCreator(
		worker_creator.Tasks(
			category_generator.CategoryGenerator(&category), location))

	for _, err := range errs {
		//noinspection GoUnhandledErrorResult
		fmt.Fprint(os.Stderr, err)
	}

	data := make(map[string]uint64)
	for _, statistic := range results {
		data[statistic.Location] = statistic.NumberOfGoods
	}

	resultJson, _ := statistics_format.StatisticsFormat{
		Location:           location,
		DataCollectionTime: time.Now().Format(time.RFC3339),
		Data:               data,
	}.MarshalJSON()

	jsonPath := "categories" + strings.Title(location) + ".json"
	file, err = os.Create(jsonPath)
	if err != nil {
		log.Fatal(err)
	}

	//noinspection GoUnhandledErrorResult
	fmt.Fprintf(os.Stderr, "The result is recorded into '"+jsonPath+"' .\n")

	//noinspection GoUnhandledErrorResult
	file.Write(resultJson)

	fmt.Print(string(resultJson))
}
