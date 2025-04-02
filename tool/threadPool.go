package tool

import "sync"

var (
	maxConcurrency int           // 最大并发数
	sem            chan struct{} //信号量
)

func NewThreadPool(maxConcurrencyVal int) {
	maxConcurrency = maxConcurrencyVal
	sem = make(chan struct{}, maxConcurrency)
}

func AppendJob(job func(), wg *sync.WaitGroup) {
	wg.Add(1)
	sem <- struct{}{} // 获取信号量
	go func() {
		defer func() {
			<-sem // 释放信号量
			wg.Done()
		}()
		job()
	}()
}
