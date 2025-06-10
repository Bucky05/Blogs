package queries

func CreateUser() string {
	return `INSERT INTO users (id, username, email, password, name) VALUES (?,?,?,?,?)`
}
func GetPasswordByUsername() string {
	return `SELECT password FROM users WHERE username = ?`
}
