package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type httpHandler struct {
}

func main() {
	fmt.Println("基于 errgroup 实现一个 http server 的启动和关闭，",
		"以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。")
	group, ctx := errgroup.WithContext(context.Background())

	//e := 0

	group.Go(func() error {
		e := signalOut()
		select {
		case <-ctx.Done():
			fmt.Println("main signalOut Done", e)
			return ctx.Err()
		}
	})

	server := http1()

	group.Go(func() error {
		fmt.Println("start listen")

		if err := server.ListenAndServe(); err != nil {
			fmt.Println("after block http", err)
			return err
		}
		select {
		case <-ctx.Done():
			if err := server.Shutdown(context.Background()); err == nil {
				fmt.Println("stop callRpc")
				fmt.Println("Shutdown http")
			}
			fmt.Println("stop")
			return ctx.Err()
		default:
			return nil
		}
	})

	group.Go(func() error {
		e := signalOut()
		if e != 0 {
			if err := server.Shutdown(context.Background()); err == nil {
				fmt.Println("stop callRpc")
				fmt.Println("Shutdown http")
			} else {
				return err
			}
		}

		select {
		case <-ctx.Done():
			fmt.Println("stop finished")
			return ctx.Err()
		default:
			return nil
		}
	})

	err := group.Wait()
	if err != nil {
		fmt.Println("errgroup wait frint:", err)
	}

}

func http1() *http.Server {
	server := &http.Server{
		Addr:           ":8808",
		Handler:        new(httpHandler),
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}

	return server
}

func (h *httpHandler) ServeHTTP(response http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello! ServerHTTP is listening.")
	response.Write([]byte("test"))
	callRpc()
}

func dbGet() {
	fmt.Println("DB Select")
}

func callRpc() {
	fmt.Println("Call RPC")
}

func signalOut() int {

	ch := make(chan os.Signal, 1)
	done := make(chan int, 1)

	defer func() {
		fmt.Println("stop dbGet")
		//close(ch)
		//close(done)
		signal.Stop(ch)
		if recover() != nil {
			fmt.Println("defer stop dbGet")
			//close(ch)
			//close(done)
			signal.Stop(ch)
		}
	}()

	// 注册
	signal.Notify(ch, os.Interrupt)

	go dbGet()

	go func() {
		signs := <-ch
		fmt.Println("sign", signs)
		done <- 1
	}()

	//fmt.Println("dbGet")
	//panic("panic sign")

	return <-done
}
