package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/*
Воспользовавшись
https://www.owasp.org/index.php/OWASP_Zed_Attack_Proxy_Project
Нашёл Url, по которому делается запрос на поиск локации.

# https://www.avito.ru/web/1/slocations?limit=1000&q=%D0%B1

Где
q - искомая подстрока, часть названия места
limit - количество вариантов выбора, <= 1000

Так как нет способа задать offset, то просто перебираем все начальные буквы.
Я проверил гипотезу кодом на Python - такой перебор вытаскивает всё дерево, не остаётся полей parent, которые бы не соответствовали какому-либо полю id.

И надо быть осторожным - всего 2-5 раз запуска утилиты, и вас могут забанить (предположительно на час).
*/

func bigramsGenerator() <-chan string {
	results := make(chan string, 50)

	go func() {
		rusAbc := "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
		for _, letter0 := range rusAbc {
			for _, letter1 := range rusAbc {
				results <- string([]rune{letter0, letter1})
			}
		}
		close(results)
	}()

	return results
}

// region ErrorForBigram
type ErrorForBigram struct {
	Bigram string
	Cause  error
}

var _ error = (*ErrorForBigram)(nil) // ErrorForBigram implement error

func (e *ErrorForBigram) Error() string {
	return "for bigram '" + e.Bigram + "': " + e.Cause.Error()
}

// endregion

func requestWaiter(bigram string, responses chan<- *LocationResponse, errors chan<- error) {

	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	resp, err := client.Get("https://www.avito.ru/web/1/slocations?limit=1000&q=" + bigram)
	if err != nil {
		errors <- &ErrorForBigram{Bigram: bigram, Cause: err}
		return
	}

	//noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	rowBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors <- &ErrorForBigram{Bigram: bigram, Cause: err}
		return
	}

	var locationResponse LocationResponse
	err = locationResponse.UnmarshalJSON(rowBody)
	if err != nil {
		errors <- &ErrorForBigram{Bigram: bigram, Cause: err}
		return
	}

	responses <- &locationResponse
}

type UniqueLocationsStorage struct {
	storage map[int64]Location
}

func NewUniqueLocationsStorage() (u *UniqueLocationsStorage) {
	return &UniqueLocationsStorage{storage: make(map[int64]Location)}
}

func (u *UniqueLocationsStorage) Add(location Location) {
	u.storage[location.Id] = location
}

func (u *UniqueLocationsStorage) GetUniques() (locations Locations) {
	for _, location := range u.storage {
		locations = append(locations, location)
	}
	return
}

func main() {
	responses := make(chan *LocationResponse, 100)
	errors := make(chan error, 100)
	var incompleteRequests uint

	for bigram := range bigramsGenerator() {
		incompleteRequests++
		go requestWaiter(bigram, responses, errors)
	}

	uniqueLocationsStorage := NewUniqueLocationsStorage()
	for incompleteRequests > 0 {
		select {
		case locationResponse := <-responses:
			incompleteRequests--
			for _, location := range locationResponse.Result.Locations {
				uniqueLocationsStorage.Add(location)
			}

		case err := <-errors:
			incompleteRequests--
			//noinspection GoUnhandledErrorResult
			fmt.Fprintln(os.Stderr, err)
		}
	}
	resultJson, err := uniqueLocationsStorage.GetUniques().MarshalJSON()
	if err != nil {
		log.Fatal()
	}

	locationsTreeFile, err := os.Create("locationsTree.json")

	//noinspection GoUnhandledErrorResult
	fmt.Fprintln(os.Stderr, "The result is recorded into './locationsTree.json' .")

	//noinspection GoUnhandledErrorResult,GoNilness
	locationsTreeFile.Write(resultJson)
	fmt.Print(string(resultJson))
}
