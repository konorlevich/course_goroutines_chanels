package task_2

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// timeoutLimit - вероятность, с которой не будет возвращаться ошибка от fakeDownload():
//
//	timeoutLimit = 100 - ошибок не будет;
//	timeoutLImit = 0 - всегда будет возвращаться ошибка.
//
// Можете "поиграть" с этим параметром, для проверки случаев с возвращением ошибки.
var timeoutLimit = 90

type Result struct {
	msg string
	err error
}

// fakeDownload - имитирует разное время скачивания для разных адресов
func fakeDownload(url string) Result {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	if r > timeoutLimit {
		return Result{
			err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url)),
		}
	}

	return Result{
		msg: fmt.Sprintf("downloaded data from %s\n", url),
	}
}

// download - параллельно скачивает данные из urls
//
// - receives an url list
// - downloads data from all the urls concurrently (using fakeDownload)
// - if fakeDownload returned and error, you need to return them all (see errors.Join)
func download(urls []string) (res []string, err error) {
	res = make([]string, 0, len(urls))
	resCh := make(chan Result)
	wg := &sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resCh <- fakeDownload(url)
		}(url)
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for s := range resCh {
		if s.err != nil {
			err = errors.Join(err, s.err)
			continue
		}
		res = append(res, s.msg)
	}

	return
}
