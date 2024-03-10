package delete_deletepayment

import (
	"ecomsvc/internal/application/account/deleteaccount"
)

type requestBody struct {
	ID int `json:"id"`
}

func (r *requestBody) toUsecasePayload() *deleteaccount.Payload {
	return &deleteaccount.Payload{
		ID: r.ID,
	}
}

type responseBody struct {
	ID      int `json:"id,omitempty"`
	UserId  int `json:"user_id,omitempty"`
	Balance int `json:"balance,omitempty"`
}

func responseFromResult(r *deleteaccount.Result) *responseBody {
	return &responseBody{
		ID:      r.ID,
		UserId:  r.UserId,
		Balance: r.Balance,
	}
}
