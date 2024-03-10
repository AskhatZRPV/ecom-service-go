package shoppingsession

type ShoppingSession struct {
	ID         int
	UserID     int
	TotalPrice int
}

func New(userID int) *ShoppingSession {
	return &ShoppingSession{
		UserID:     userID,
		TotalPrice: 0,
	}
}

func (s *ShoppingSession) UpdatePrice(newPrice int) *ShoppingSession {
	return &ShoppingSession{
		ID:         s.ID,
		UserID:     s.ID,
		TotalPrice: newPrice,
	}
}
