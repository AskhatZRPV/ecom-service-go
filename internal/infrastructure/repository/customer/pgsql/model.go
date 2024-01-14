package pgsql

type customerRow struct {
	Id        int
	UserId    int
	FirstName string
	LastName  string
	Address   string
	City      string
	State     string
	ZIPCode   string
	Country   string
}
