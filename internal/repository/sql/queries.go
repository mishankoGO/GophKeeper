package sql

var RegisterQuery = `INSERT INTO credentials ("login", "hash_password") VALUES ($1, $2)`
var GetUserId = `SELECT user_id FROM credentials WHERE login = $1`
var AddUserQuery = `INSERT INTO users (user_id, login, created_at) VALUES ($1, $2, $3)`
var CheckUser = `SELECT COUNT(*) FROM credentials WHERE login = $1`
var LoginUser = `SELECT c.user_id, c.hash_password, u.created_at FROM credentials c JOIN users u ON u.user_id = c.user_id WHERE c.login = $1`

var InsertBinaryFile = `INSERT INTO binary_files (user_id, name, hash_file, updated_at, meta) VALUES ($1, $2, $3, $4, $5)`

var GetBinaryFile = `SELECT user_id, name, hash_file, updated_at, meta FROM binary_files WHERE user_id = $1 AND name = $2`

var UpdateBinaryFile = `UPDATE binary_files SET user_id=$1, name=$2, hash_file=$3, updated_at=$4, meta=$5 WHERE user_id = $1`

var DeleteBinaryFile = `DELETE FROM binary_files WHERE user_id = $1 AND name = $2`
