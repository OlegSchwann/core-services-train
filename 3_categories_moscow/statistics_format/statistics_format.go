package statistics_format

// тип для записи конечного результата в файл

// easyjson:json
type StatisticsFormat struct {
	Location           string            `json:"location"`
	DataCollectionTime string            `json:"data_collection_time"` // время в формате RFC 3339
	Data               map[string]uint64 `json:"data"`
}
