package router

import (
	"github.com/gin-gonic/gin"
	"githubnote/kubernetes-dev/kubernetes-operator/sample/day2get-web/pkg/apis"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", apis.Ping)
	r.GET("/namespaces", apis.GetNamespaces)
}
