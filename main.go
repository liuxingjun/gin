package main

import (
	"context"
	"gin/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init()  {
	logFile, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	gin.ForceConsoleColor()
	if gin.Mode() == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}
}
func main() {

	engine := gin.Default()
	// Disable Console Color
	// gin.DisableConsoleColor()

	//r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	// your custom format
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC3339),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		//param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	//r.Use(gin.Recovery())

	router.SetupRouter(engine)

	// Listen and Server in 0.0.0.0:8080

	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Printf("listen:%s",server.Addr)
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	//engine.Run(":8080")
}
