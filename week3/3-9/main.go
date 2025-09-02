package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

/**
为什么要启动一个独立的 goroutine 来运行 HTTP 服务器？
在 Gin 或任何 Go HTTP 服务器中，将服务器启动放在独立的 goroutine 中运行有几个重要原因：

1. 避免阻塞主 goroutine
go
// 如果直接调用（不在goroutine中），这行代码会阻塞
r.Run(":9098")

// 后面的信号监听代码将永远无法执行
exit := make(chan os.Signal)
signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
<-exit // 这行代码永远不会到达
HTTP 服务器的 Run 或 ListenAndServe 方法是阻塞调用，它会一直运行直到服务器关闭。如果不在单独的 goroutine 中运行，它会阻止后续代码的执行。

2. 允许主 goroutine 处理其他任务
通过将服务器放在独立的 goroutine 中，主 goroutine 可以自由地处理其他任务，例如：

监听系统信号以实现优雅关闭

执行健康检查或监控

处理其他并发任务

管理应用程序的生命周期

3. 实现优雅关闭机制
这是最重要的原因之一。通过分离服务器 goroutine 和主控制 goroutine，可以实现：

go
func main() {
    // ... 服务器设置 ...

    // 服务器在独立goroutine中运行
    go func() {
        srv.ListenAndServe()
    }()

    // 主goroutine可以监听退出信号
    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit // 等待退出信号

    // 收到信号后，可以优雅地关闭服务器
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    srv.Shutdown(ctx) // 优雅关闭
}
4. 符合 Go 的并发模型
Go 的并发模型鼓励使用 goroutine 来处理阻塞 I/O 操作。HTTP 服务器本质上是一个 I/O 密集型任务，非常适合在独立的 goroutine 中运行。

5. 支持多个服务器实例
在某些场景下，您可能需要在同一程序中运行多个服务器：

go
func main() {
    // API 服务器
    go func() {
        apiServer := gin.Default()
        apiServer.Run(":8080")
    }()

    // 监控服务器
    go func() {
        metricsServer := gin.Default()
        metricsServer.Run(":9090")
    }()

    // 主goroutine处理信号和协调
    // ...
}
6. 保持代码清晰和可维护性
分离关注点是良好的软件设计实践。将服务器运行逻辑与应用程序控制逻辑分离，使代码更加模块化和易于维护。

总结
启动独立 goroutine 来运行 HTTP 服务器的主要目的是避免阻塞主控制流，从而允许主 goroutine 处理其他重要任务，特别是实现优雅关闭机制。
这种模式是 Go Web 开发中的标准做法，确保了应用程序可以正确地响应系统信号并优雅地处理关闭过程。
*/

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	go func() {
		r.Run(":9098")
	}()
	exit := make(chan os.Signal)
	//syscall.SIGINT 是操作系统信号的一种，表示"中断信号"(Interrupt Signal)。
	// 通知 signal 包将 SIGINT 和 SIGTERM 信号转发到 exit 通道
	//syscall.SIGINT 是表示用户中断请求的系统信号，通常由 Ctrl+C 触发。
	//在 Go 程序中捕获和处理此信号允许实现优雅退出机制，确保程序在终止前能够完成必要的清理工作，
	//这对于生产环境中的应用程序至关重要。
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	//等待通道信息，然后回继续执行
	// 阻塞等待，直到收到上述任一信号
	<-exit
	fmt.Println("程序优雅退出，结束...")
}
