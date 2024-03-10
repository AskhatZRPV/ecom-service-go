package post_createpayment

import (
	"ecomsvc/internal/application/account/createaccount"
)

type requestBody struct {
	UserId  int `json:"user_id,omitempty"`
	Balance int `json:"balance,omitempty"`
}

func (r *requestBody) toUsecasePayload() *createaccount.Payload {
	return &createaccount.Payload{
		UserId:  r.UserId,
		Balance: r.Balance,
	}
}

type responseBody struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	AccountId int    `json:"account_id"`
}

func responseFromResult(id int) *responseBody {
	return &responseBody{
		Status:    "Success",
		Message:   "Successfully created account",
		AccountId: id,
	}
}
