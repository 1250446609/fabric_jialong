package web

import (
	"encoding/json"
	"fmt"
	"github.com/kongyixueyuan.com/alarm/web/controller"
	"net/http"
)

// 启动Web服务并指定路由信息
func WebStart(app controller.Application) {

	http.HandleFunc("/addAlarm", app.AddAlarm) // 提交信息请求

	fmt.Println("启动Web服务, 监听端口号为: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Web服务启动失败: %v", err)
	}
}
