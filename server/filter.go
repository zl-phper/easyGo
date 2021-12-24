package server

import (
	"fmt"
	"time"
)

type handleFunc func(c *Context)

type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)

var  _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {

	return func(c *Context) {

		start := time.Now().Nanosecond()

		next(c)

		end := time.Now().Nanosecond()

		fmt.Printf("用了 %d 纳秒", end - start)
	}
}
func MetricsFilterBuilder2(next Filter) Filter {

	return func(c *Context) {

		go func() {
		 	time.Sleep(time.Second * 10)

 			fmt.Printf("这里执行了异步请求")
		}()



	}
}
