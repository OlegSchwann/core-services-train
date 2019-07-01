package worker_creator

import (
	"github.com/OlegSchwann/core-services-train/3_categories_moscow/worker"
	"time"
)

// Task - задача для системы скачивания: для какого города,
// для какой категории взять данные и сколько осталось попыток.
type Task struct {
	Location string
	Category string
}

// DownloadedStatistics скаченное число товаров на странице.
type DownloadedStatistics struct {
	Task
	NumberOfGoods uint64
}

// Сообщение о ошибке, если не удалось скачать.
type DownloadErr struct {
	Task  // Task.TimeToLive уменьшено на 1.
	Cause error
}

func Tasks(categorys []string, location string) (tasks []Task) {
	tasks = make([]Task, len(categorys))
	for i, category := range categorys {
		tasks[i].Location = location
		tasks[i].Category = category
	}
	return
}

// WorkerCreator Планировщик для обработчиков запросов.
// Запускает задачи на скачивание.
// Завершается, если все задачи скачаны, или лимит повторных попыток исчерпан.
func WorkerCreator(tasks []Task) (results []DownloadedStatistics, downloadErrs []DownloadErr) {

	_results := make(chan DownloadedStatistics, 20)
	_downloadErrs := make(chan DownloadErr, 20)

	for _, task := range tasks {
		go func(task Task) {
			var err error
			for attempts := 2; attempts != 0; attempts-- {
				var numberOfGoods uint64
				numberOfGoods, err = worker.Worker(task.Location, task.Category)
				if err != nil {
					// Вдруг это мгновенная неполадка сети?
					time.Sleep(500 * time.Microsecond)
					continue
				}
				_results <- DownloadedStatistics{
					Task:          task,
					NumberOfGoods: numberOfGoods,
				}
				return
			}
			_downloadErrs <- DownloadErr{
				Task:  task,
				Cause: err,
			}
			return
		}(task)
	}

	for i := 0; i != len(tasks); i++ {
		select {
		case result := <-_results:
			results = append(results, result)
		case downloadErr := <-_downloadErrs:
			downloadErrs = append(downloadErrs, downloadErr)
		}
	}
	return
}
