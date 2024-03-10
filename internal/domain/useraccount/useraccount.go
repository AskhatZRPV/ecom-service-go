package useraccount

type UserAccount struct {
	ID      int
	UserId  int
	Balance int
}

func New(userId int, balance int) *UserAccount {
	return &UserAccount{
		UserId:  userId,
		Balance: balance,
	}
}

func (u *UserAccount) Update(balance int) *UserAccount {
	return &UserAccount{
		ID:      u.ID,
		UserId:  u.UserId,
		Balance: balance,
	}
}

func Update(id int, balance int) *UserAccount {
	return &UserAccount{
		ID:      id,
		Balance: balance,
	}
}
