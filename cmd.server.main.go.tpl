package main

import "github.com/gin-gonic/gin"

func main() {
	s, err := NewServices()
	if err != nil {
		panic(err)
	}

	e := gin.New()
	s.register(e)

	svr := &http.Server{
		Addr:           ":8090",
		Handler:        e,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}

	go func() {
		err := svr.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("svr.ListenAndServe err: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal)
	// 接受 syscall.SIGINT 和 syscall.SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	// 最大时间控制，用于通知该服务端它有 5 秒的时间处理原来的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server extting")
}