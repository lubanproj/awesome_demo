package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var concurrencyV2 = flag.Int64("concurrency", 20, "concurrency")
var totalV2 = flag.Int64("total", 2000, "total requests")

func main() {
	flag.Parse()
	requestV2(*totalV2, *concurrencyV2)
}


func requestV2(totalReqs int64, concurrency int64) {

	perClientReqs := totalReqs / concurrency

	counter := &CounterV2{
		Total: perClientReqs * concurrency ,
		Concurrency: concurrency,
	}

	var wg sync.WaitGroup
	wg.Add(int(concurrency))

	startTime := time.Now().UnixNano()

	for i:=int64(0); i<counter.Concurrency; i++ {

		go func(i int64) {


			for j:=int64(0); j< perClientReqs; j++ {

				rsp, err := callV3()
				if err != nil {
					fmt.Printf("could not greet: %v\n", err)
				}

				if err == nil && rsp == "world" {
					atomic.AddInt64(&counter.Succ, 1)
				} else {
					fmt.Printf("rsp fail : %v\n", err)
					atomic.AddInt64(&counter.Fail, 1)
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()

	counter.Cost = (time.Now().UnixNano() - startTime) / 1000000

	fmt.Printf("took %d ms for %d requests \n", counter.Cost, counter.Total)
	fmt.Printf("sent     requests      : %d\n", counter.Total)
	fmt.Printf("received requests      : %d\n", atomic.LoadInt64(&counter.Succ) + atomic.LoadInt64(&counter.Fail))
	fmt.Printf("received requests succ : %d\n", atomic.LoadInt64(&counter.Succ))
	fmt.Printf("received requests fail : %d\n", atomic.LoadInt64(&counter.Fail))
	fmt.Printf("throughput  (TPS)      : %d\n", totalReqs*1000/counter.Cost)

}

type CounterV2 struct {
	Succ int64  // 成功量
	Fail int64  // 失败量
	Total int64 // 总量
	Concurrency int64 // 并发量
	Cost int64  // 总耗时 ms
}


