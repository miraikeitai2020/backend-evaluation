package model

/*--------------------------------*/
/*     Handler Request Struct     */
/*--------------------------------*/
type EvaluateSpotRequest struct {
	ID		string	`json:"id"`
	Emotion	int		`json:"emotion"`
	Status	bool	`json:"status"`
}
