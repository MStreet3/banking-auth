package domain

type User struct {
	Username   string `db:"username"`
	Password   string `db:"username"`
	Role       string `db:"role"`
	CustomerId string `db:"customer_id"`
}
