package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"githubnote/kubernetes-dev/kubernetes-operator/sample/day2get-web/pkg/config"
	"githubnote/kubernetes-dev/kubernetes-operator/sample/day2get-web/pkg/router"
	"k8s.io/klog/v2"
)

func main() {
	engine := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.InitRouter(engine)
	err := engine.Run(fmt.Sprintf("%s:%d", config.GetString(config.ServerHost), config.GetInt(config.ServerPort)))
	if err != nil {
		klog.Fatal(err)
		return
	}
}
