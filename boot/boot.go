package boot

import (
	"context"
	"fmt"
	"gsadmin/core/config"
	"gsadmin/core/db"
	"gsadmin/core/log"
	"gsadmin/core/queue"
	"gsadmin/core/utils/assertion"
	"gsadmin/core/utils/translator"
	"gsadmin/database"
	"gsadmin/global/e"
	"gsadmin/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitServer() { //staticFs, templateFs embed.FS
	//初始化翻译器
	_ = translator.InitTrans("zh")

	//初始化配置文件
	config.InitConfig("./config.toml")
	//初始化日志
	log.InitLog()
	//初始化数据库
	db.InitConn()
	database.InitTables()

	//初始化消息队列
	registerQueue()

	//初始化路由
	r := initRouter() //staticFs, templateFs
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Instance().App.HttpPort),
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//欢迎语
	usage()
	//启动服务
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Instance().Error(err.Error())
			os.Exit(0)
		}
	}()
	//退出
	shutDown(s)
}

func usage() {
	usageStr := fmt.Sprintf(`	欢迎使用 %s@%s
	程序运行地址:http://127.0.0.1:%s...`, config.Instance().App.Name, config.Instance().App.Version, assertion.AnyToString(config.Instance().App.HttpPort))
	log.Instance().Info("Start server, running at: http://127.0.0.1:" + assertion.AnyToString(config.Instance().App.HttpPort))
	fmt.Printf("%s\n", usageStr)
}

func shutDown(s *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Instance().Info("Server exiting.")
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := db.Instance().Close()
	if err != nil {
		log.Instance().Fatal("Close DB error: " + err.Error())
	}

	if err = s.Shutdown(ctx); err != nil {
		log.Instance().Fatal("Shutdown server error: " + err.Error())
	}
	fmt.Println("Server exited.")
}

func registerQueue() {
	//注册日志队列
	_ = queue.Instance().RegisterTopic(e.TopicOperLog)
	go queue.Instance().Subscribe(e.TopicOperLog, middleware.WriteTo)
	//注册业务队列
	//TODO
}
