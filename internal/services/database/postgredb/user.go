package postgredb

type User struct {
	Id        int
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func (hand *Postgres) ReadUser(userName string) (usr *User, err error) {
	row := hand.db.QueryRow("SELECT * FROM users WHERE username = $1", userName)
	usr = new(User)
	err = row.Scan(&usr.Id, &usr.Username, &usr.Password, &usr.FirstName, &usr.LastName)
	return
}
