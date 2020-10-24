package view

/*--------------------------------------*/
/*      Handler Response Generator      */
/*--------------------------------------*/
type MutationResponse struct {
	Status	bool	`json:"status"`
	Error	Error	`json:"error"`
}

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}

func MakeMutationResponse(status bool, code int, value ...string) (body MutationResponse) {
	body.Status = status
	if !status {
		body.Error = Error{
			Code: code,
			Message: value[0],
			Description: value[1],
		}
	}
	return
}
