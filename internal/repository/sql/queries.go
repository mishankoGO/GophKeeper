package sql

// Registration and login.
var RegisterQuery = `INSERT INTO credentials (login, password) VALUES ($1, $2)`
var GetUserId = `SELECT user_id FROM credentials WHERE login = $1`
var AddUserQuery = `INSERT INTO users (user_id, login, created_at) VALUES ($1, $2, $3)`
var CheckUser = `SELECT COUNT(*) FROM credentials WHERE login = $1`
var LoginUser = `SELECT c.user_id, c.password, u.created_at FROM credentials c JOIN users u ON u.user_id = c.user_id WHERE c.login = $1`

// Binary file queries.
var InsertBinaryFile = `INSERT INTO binary_files (user_id, name, file, updated_at, meta) VALUES ($1, $2, $3, $4, $5)`

var GetBinaryFile = `SELECT user_id, name, file, updated_at, meta FROM binary_files WHERE user_id = $1 AND name = $2`

var UpdateBinaryFile = `UPDATE binary_files SET user_id=$1, name=$2, file=$3, updated_at=$4, meta=$5 WHERE user_id = $1 AND name=$2`

var DeleteBinaryFile = `DELETE FROM binary_files WHERE user_id = $1 AND name = $2`

var ListBinaryFiles = `SELECT name, file, updated_at, meta FROM binary_files WHERE user_id = $1`

// Cards queries.
var InsertCard = `INSERT INTO cards (user_id, name, card, updated_at, meta) VALUES ($1, $2, $3, $4, $5)`

var GetCard = `SELECT user_id, name, card, updated_at, meta FROM cards WHERE user_id = $1 AND name = $2`

var UpdateCard = `UPDATE cards SET user_id=$1, name=$2, card=$3, updated_at=$4, meta=$5 WHERE user_id = $1 AND name=$2`

var DeleteCard = `DELETE FROM cards WHERE user_id = $1 AND name = $2`

var ListCards = `SELECT name, card, updated_at, meta FROM cards WHERE user_id = $1`

// Log pass queries
var InsertLogPass = `INSERT INTO log_passes (user_id, name, login, password, updated_at, meta) VALUES ($1, $2, $3, $4, $5, $6)`

var GetLogPass = `SELECT user_id, name, login, password, updated_at, meta FROM log_passes WHERE user_id = $1 AND name = $2`

var UpdateLogPass = `UPDATE log_passes SET user_id=$1, name=$2, login=$3, password=$4, updated_at=$5, meta=$6 WHERE user_id = $1 AND name=$2`

var DeleteLogPass = `DELETE FROM log_passes WHERE user_id = $1 AND name = $2`

var ListLogPasses = `SELECT name, login, password, updated_at, meta FROM log_passes WHERE user_id = $1`

// Text queries
var InsertText = `INSERT INTO texts (user_id, name, text, updated_at, meta) VALUES ($1, $2, $3, $4, $5)`

var GetText = `SELECT user_id, name, text, updated_at, meta FROM texts WHERE user_id = $1 AND name = $2`

var UpdateText = `UPDATE texts SET user_id=$1, name=$2, text=$3, updated_at=$4, meta=$5 WHERE user_id = $1 AND name=$2`

var DeleteText = `DELETE FROM texts WHERE user_id = $1 AND name = $2`

var ListTexts = `SELECT name, text, updated_at, meta FROM texts WHERE user_id = $1`
