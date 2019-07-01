package worker

import (
	"testing"
)

func TestWorker(t *testing.T) {
	numberOfGoods, err := Worker("sankt-peterburg", "/mototsikly_i_mototehnika/mototsikly")
	if err != nil {
		t.Error("expected, uint nil, got err:\n", err)
	}
	t.Log("on https://www.avito.ru/sankt-peterburg/mototsikly_i_mototehnika/mototsikly found ", numberOfGoods, "product")
}
