package controller

import(
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/miraikeitai2020/backend-evaluation/pkg/database"
	"github.com/miraikeitai2020/backend-evaluation/pkg/server/view"
	"github.com/miraikeitai2020/backend-evaluation/pkg/server/model"
)

type Controller struct {
	DB	*database.DB
}

func Init(db *database.DB) Controller {
	return Controller{
		DB: db,
	}
}

func (ctrl *Controller) EvaluateSpot(cxt *gin.Context) {
	var request model.EvaluateSpotRequest
	err := cxt.BindJSON(&request)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		cxt.JSON(
			http.StatusInternalServerError,
			view.MakeMutationResponse(
				false,
				http.StatusInternalServerError,
				"Internal Server Error",
				fmt.Sprintf("%v", err),
			),
		)
	}
	ctrl.DB.UpdateSpotEmotionValue(request.ID, request.Emotion, request.Status)
	cxt.JSON(
		http.StatusOK,
		view.MakeMutationResponse(
			true,
			http.StatusOK,
		),
	)
}
