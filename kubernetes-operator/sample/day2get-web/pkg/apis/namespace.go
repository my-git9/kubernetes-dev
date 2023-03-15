package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"githubnote/kubernetes-dev/kubernetes-operator/sample/day2get-web/pkg/service"
)

func GetNamespaces(c *gin.Context) {
	namespace, err := service.GetNamespaes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, namespace)
}
