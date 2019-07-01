package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var pattern = regexp.MustCompile(
	`<span class="page-title-count" data-marker="page-title/count">([ \d]+)</span>`)

func Worker(location string, productCategory string) (numberOfGoods uint64, err error) {
	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	url := "https://www.avito.ru/" + location + productCategory

	resp, err := client.Get(url)
	if err != nil {
		err = errors.Wrap(err, "unable to GET '"+url+"', maybe you were banned on the site?")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "unable to read ioutil.ReadAll(resp.Body): ")
		return
	}

	result := pattern.FindStringSubmatch(string(body))
	if len(result) < 2 {
		err = errors.New(fmt.Sprint("unable to find numberOfGoods on "+url+": got only ", result, "\n\nRequest body:\n\n", string(body), "are you baned?\n\n"))
		return
	}

	numberOfGoods, err = strconv.ParseUint(strings.Replace(result[1], " ", "", -1), 10, 64)
	return
}
