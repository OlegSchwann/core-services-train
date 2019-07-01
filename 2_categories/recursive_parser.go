package main

import (
	"avitoTest/2_categories/category_item"
	"golang.org/x/net/html"
)

func titleHrefParser(div *html.Node) (title string, href string) {
	/*
		div
		   	i (skip)
		   	a title href
	*/
	for a := div.FirstChild; a != nil; a = a.NextSibling {
		if a.Type != html.ElementNode && a.Data != "a" {
			continue
		}
		for _, attr := range a.Attr {
			switch attr.Key {
			case "href":
				href = "https://www.avito.ru" + attr.Val
			case "title":
				title = attr.Val
			}
		}
	}
	return
}

func recursiveParser(ui *html.Node) (categoryItems category_item.CategoryItems) {
	/*
		ui
		   	li
		   		div
		   			i (skip)
		   			a title href
		   		ui (recursive call)
	*/
	for li := ui.FirstChild; li != nil; li = li.NextSibling {
		if li.Type != html.ElementNode || li.Data != "li" /* может быть "button" */ {
			continue
		}

		var categoryItem category_item.CategoryItem

		for divOrUi := li.FirstChild; divOrUi != nil; divOrUi = divOrUi.NextSibling {
			if divOrUi.Type != html.ElementNode {
				continue
			}

			switch divOrUi.Data {
			case "div":
				categoryItem.Title, categoryItem.Href = titleHrefParser(divOrUi)
			case "ul":
				categoryItem.Childes = recursiveParser(divOrUi)
			}
		}

		categoryItems = append(categoryItems, categoryItem)
	}
	return
}
