/*
	基于 errgroup 实现一个 http server 的启动和关闭 ，
	以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/
package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Sqc"))
	})
	server := http.Server{
		Handler: mux,
		Addr:    "127.0.0.1:8000",
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})
	shutdown := make(chan struct{})
	g.Go(func() error {
		<-shutdown
		log.Println("shutting down server...")
		tCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		err := server.Shutdown(tCtx)

		return err
	})

	g.Go(func() error {
		return signalHandle(shutdown)
	})

	log.Printf("%+v end", g.Wait())
}

func signalHandle(shutdown chan<- struct{}) error {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	shutdown <- struct{}{}
	return errors.New("signal shutdown")
}
