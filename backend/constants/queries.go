package constants

var QueryMap = map[string]string{
	"createUser": `INSERT INTO users (id, email, password, name) VALUES (?,?,?,?)`,
}
