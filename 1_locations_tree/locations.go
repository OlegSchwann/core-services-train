package main

/*
Формат приходящих данных
https://www.avito.ru/web/1/slocations?limit=1000&q=%D0%B1

{
  'result': {
    'locations': [
      {
        'id': 623840,
        'names': {
          '1': 'Брянская область'
        }
      },
      {
        'id': 621630,
        'parent': {
          'id': 621590,
          'names': {
            '1': 'Алтайский край'
          }
        },
        'names': {
          '1': 'Барнаул'
        }
      }
    ]
  }
}

*/

// Соответствующая структура данных для разворачивания.

//easyjson:json
type LocationResponse struct {
	Result struct {
		Locations Locations `json:"locations"`
	} `json:"result"`
}

//easyjson:json
type Locations []Location

//easyjson:json
type Location struct {
	Id    int64 `json:"id"`
	Names struct {
		First string `json:"1"`
	} `json:"names"`
	Parent *struct {
		Id    int64 `json:"id"`
		Names struct {
			First string `json:"1"`
		} `json:"names"`
	} `json:"parent,optional,omitempty"`
}
