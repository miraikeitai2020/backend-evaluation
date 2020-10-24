package server

import(
	"github.com/gin-gonic/gin"
	"github.com/miraikeitai2020/backend-evaluation/pkg/server/controller"
)

func Router(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	router.POST("/mutation/evaluate/spot", ctrl.EvaluateSpot)
	return router
}
