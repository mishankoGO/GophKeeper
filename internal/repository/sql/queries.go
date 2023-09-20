package sql

var RegisterQuery = `INSERT INTO credentials ("login", "hash_password") VALUES ($1, $2)`
var GetUserId = `SELECT user_id FROM credentials WHERE login = $1`
var AddUserQuery = `INSERT INTO users ("user_id", "created_at") VALUES ($1, $2)`
