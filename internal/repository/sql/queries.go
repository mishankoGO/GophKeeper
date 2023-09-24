package sql

var RegisterQuery = `INSERT INTO credentials ("login", "hash_password") VALUES ($1, $2)`
var GetUserId = `SELECT user_id FROM credentials WHERE login = $1`
var AddUserQuery = `INSERT INTO users (user_id, login, created_at) VALUES ($1, $2, $3)`
var CheckUser = `SELECT COUNT(*) FROM credentials WHERE login = $1`
var LoginUser = `SELECT c.user_id, c.hash_password, u.created_at FROM credentials c JOIN users u ON u.user_id = c.user_id WHERE c.login = $1`
