package category_generator

import (
	"avitoTest/2_categories/category_item"
	"regexp"
)

var urlExtractLocation = regexp.MustCompile(`https://www\.avito\.ru/rossiya([/\w]+)\??.*`)

// Рекурсивно обходит дерево локаций.
func CategoryGenerator(tree *category_item.CategoryItem) (locations []string) {
	var searcher func(items category_item.CategoryItems)
	searcher = func(items category_item.CategoryItems) {
		// Базовый случай - выходим, если нет детей.
		for _, item := range items {
			// Рекурсивный случай - для каждого ребёнка печатаем вырезанную из href часть с категорией.
			locations = append(locations, urlExtractLocation.FindStringSubmatch(item.Href)[1])
			searcher(item.Childes)
		}
		// пропускаем корень - на главной странице нестандартное отображение количества товаров.
	}
	searcher(tree.Childes)

	return
}
