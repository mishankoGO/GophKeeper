package sql

// Registration and login.
var RegisterQuery = `INSERT INTO credentials ("login", "hash_password") VALUES ($1, $2)`
var GetUserId = `SELECT user_id FROM credentials WHERE login = $1`
var AddUserQuery = `INSERT INTO users (user_id, login, created_at) VALUES ($1, $2, $3)`
var CheckUser = `SELECT COUNT(*) FROM credentials WHERE login = $1`
var LoginUser = `SELECT c.user_id, c.hash_password, u.created_at FROM credentials c JOIN users u ON u.user_id = c.user_id WHERE c.login = $1`

// Binary file queries.
var InsertBinaryFile = `INSERT INTO binary_files (user_id, name, hash_file, updated_at, meta) VALUES ($1, $2, $3, $4, $5)`

var GetBinaryFile = `SELECT user_id, name, hash_file, updated_at, meta FROM binary_files WHERE user_id = $1 AND name = $2`

var UpdateBinaryFile = `UPDATE binary_files SET user_id=$1, name=$2, hash_file=$3, updated_at=$4, meta=$5 WHERE user_id = $1`

var DeleteBinaryFile = `DELETE FROM binary_files WHERE user_id = $1 AND name = $2`

// Cards queries.
var InsertCard = `INSERT INTO cards (user_id, name, hash_card_number, hash_card_holder, expiry_date, hash_cvv, updated_at, meta) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

var GetCard = `SELECT user_id, name, hash_card_number, hash_card_holder, expiry_date, hash_cvv, updated_at, meta FROM cards WHERE user_id = $1 AND name = $2`

var UpdateCard = `UPDATE cards SET user_id=$1, name=$2, hash_card_number=$3, hash_card_holder=$4, expiry_date=$5, hash_cvv=$6, updated_at=$7, meta=$8 WHERE user_id = $1`

var DeleteCard = `DELETE FROM cards WHERE user_id = $1 AND name = $2`

// Log pass queries
var InsertLogPass = `INSERT INTO log_passes (user_id, name, hash_login, hash_password, updated_at, meta) VALUES ($1, $2, $3, $4, $5, $6)`

var GetLogPass = `SELECT user_id, name, hash_login, hash_password, updated_at, meta FROM log_passes WHERE user_id = $1 AND name = $2`

var UpdateLogPass = `UPDATE log_passes SET user_id=$1, name=$2, hash_login=$3, hash_password=$4, updated_at=$7, meta=$8 WHERE user_id = $1`

var DeleteLogPass = `DELETE FROM log_passes WHERE user_id = $1 AND name = $2`
