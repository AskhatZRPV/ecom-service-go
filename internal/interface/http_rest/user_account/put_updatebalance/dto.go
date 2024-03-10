package put_updatebalance

import (
	"ecomsvc/internal/application/account/updateaccount"
)

type requestBody struct {
	ID      int `json:"id"`
	Balance int `json:"balance"`
}

func (r *requestBody) toUsecasePayload() *updateaccount.Payload {
	return &updateaccount.Payload{
		ID:      r.ID,
		Balance: r.Balance,
	}
}

type responseBody struct {
	ID      int `json:"id,omitempty"`
	UserId  int `json:"user_id,omitempty"`
	Balance int `json:"balance,omitempty"`
}

func responseFromResult(r *updateaccount.Result) *responseBody {
	return &responseBody{
		ID:      r.ID,
		UserId:  r.UserId,
		Balance: r.Balance,
	}
}
