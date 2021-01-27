package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("1. 用 Go 实现一个 tcp server，",
		"用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出")
	fmt.Println("https://github.com/Go-000/Go-000/issues/82")
	tcpAddr := net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 6666,
		Zone: "",
	}

	tcp := server(&tcpAddr)
	fmt.Println(tcp.Addr().String())
	defer func() {
		tcp.Close()
		fmt.Println("tcp Close")
	}()
	var ch = make(chan string)

	gCtx := context.Background()
	gCtx = context.WithValue(gCtx, "conn", 0)
	gCtx, cancel := context.WithTimeout(gCtx, time.Second*10)
	defer cancel()

	ctx := context.Background()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	go func(gCtx context.Context) {
		for {
			var errorCount int

			if errorCount > 2 {
				panic("errorCount > 2")
			}

			conn, err := tcp.AcceptTCP()

			fmt.Println(conn)
			if err != nil {
				log.Fatal("conn error: ", err)
			}
			defer conn.Close()

			if err != nil {
				log.Println("accept error: ", err)
				errorCount++
				time.Sleep(time.Second * 1)
			}

			go func(tcpConn *net.TCPConn) {
				var b []byte
				for can := true; can; {
					ctx, cancel := context.WithTimeout(ctx, time.Second*2)
					defer cancel()

					fmt.Println("get connect: ", tcpConn.RemoteAddr().String())
					if _, err := tcpConn.Read(b); err != nil {
						log.Println("read error: ", err)
					}
					fmt.Println(string(b))
					//ch <- string(b)
					select {
					case <-ctx.Done():
						fmt.Println("ctx Done:", ctx.Err())
						conn.Close()
						can = false
					}
				}
			}(conn)

			go func(tcpConn *net.TCPConn) {
				str := <-ch
				_, err := tcpConn.Write([]byte("Server receive succeed." + str + " Ack: OK\n"))
				if err != nil {
					log.Println("Send error:", err)
				}
			}(conn)
		}
	}(gCtx)

	select {
	case <-sig:
		fmt.Println("close")
		fmt.Println("NOW conn", gCtx.Value("conn").(int))
	}
}

func server(tcpAddr *net.TCPAddr) net.TCPListener {
	tcp, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	return *tcp
}
