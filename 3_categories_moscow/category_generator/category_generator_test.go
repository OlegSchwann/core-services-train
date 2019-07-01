package category_generator

import (
	"fmt"
	"github.com/OlegSchwann/core-services-train/2_categories/category_item"
	"reflect"
	"strconv"
	"testing"
)

func TestGenerator(t *testing.T) {
	testCategoryItem := category_item.CategoryItem{
		Title: "Бытовая электроника",
		Href:  "https://www.avito.ru/rossiya/bytovaya_elektronika?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "Аудио и видео",
			Href:  "https://www.avito.ru/rossiya/audio_i_video?s_trg=11&cd=1",
			Childes: []category_item.CategoryItem{{
				Title: "MP3-плееры",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/mp3-pleery?s_trg=11&cd=1",
			}, {
				Title: "Акустика, колонки, сабвуферы",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/akustika_kolonki_sabvufery?s_trg=11&cd=1",
			}, {
				Title: "Видео, DVD и Blu-ray плееры",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/video_dvd_i_blu-ray_pleery?s_trg=11&cd=1",
			}, {
				Title: "Видеокамеры",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/videokamery?s_trg=11&cd=1",
			}, {
				Title: "Кабели и адаптеры",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/kabeli_i_adaptery?s_trg=11&cd=1",
			}, {
				Title: "Микрофоны",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/mikrofony?s_trg=11&cd=1",
			}, {
				Title: "Музыка и фильмы",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/muzyka_i_filmy?s_trg=11&cd=1",
			}, {
				Title: "Музыкальные центры, магнитолы",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/muzykalnye_tsentry_magnitoly?s_trg=11&cd=1",
			}, {
				Title: "Наушники",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/naushniki?s_trg=11&cd=1",
			}, {
				Title: "Телевизоры и проекторы",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/televizory_i_proektory?s_trg=11&cd=1",
			}, {
				Title: "Усилители и ресиверы",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/usiliteli_i_resivery?s_trg=11&cd=1",
			}, {
				Title: "Аксессуары",
				Href:  "https://www.avito.ru/rossiya/audio_i_video/aksessuary?s_trg=11&cd=1"}},
		}, {
			Title: "Игры, приставки и программы",
			Href:  "https://www.avito.ru/rossiya/igry_pristavki_i_programmy?s_trg=11&cd=1",
			Childes: []category_item.CategoryItem{{
				Title: "Игры для приставок",
				Href:  "https://www.avito.ru/rossiya/igry_pristavki_i_programmy/igry_dlya_pristavok?s_trg=11&cd=1",
			}, {
				Title: "Игровые приставки",
				Href:  "https://www.avito.ru/rossiya/igry_pristavki_i_programmy/igrovye_pristavki?s_trg=11&cd=1",
			}},
		}},
	}

	expectedResult := []string{
		"/audio_i_video",
		"/audio_i_video/mp3",
		"/audio_i_video/akustika_kolonki_sabvufery",
		"/audio_i_video/video_dvd_i_blu",
		"/audio_i_video/videokamery",
		"/audio_i_video/kabeli_i_adaptery",
		"/audio_i_video/mikrofony",
		"/audio_i_video/muzyka_i_filmy",
		"/audio_i_video/muzykalnye_tsentry_magnitoly",
		"/audio_i_video/naushniki",
		"/audio_i_video/televizory_i_proektory",
		"/audio_i_video/usiliteli_i_resivery",
		"/audio_i_video/aksessuary",
		"/igry_pristavki_i_programmy",
		"/igry_pristavki_i_programmy/igry_dlya_pristavok",
		"/igry_pristavki_i_programmy/igrovye_pristavki"}

	var categories []string
	for category := range CategoryGenerator(&testCategoryItem) {
		categories = append(categories, strconv.Itoa(category))
	}

	if !reflect.DeepEqual(categories, expectedResult) {
		t.Error(fmt.Sprintf("Expected\n\n: %#v\n\n, got:\n\n %#v\n\n", expectedResult, categories))
	}
}
