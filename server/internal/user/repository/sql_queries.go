package repository

const (
	createUserQuery = `INSERT INTO users (login, password) 
		VALUES ($1, $2) 
		RETURNING user_id, login, password`

	findByLoginQuery = `SELECT user_id, login, password FROM users WHERE login = $1`

	findByIDQuery = `SELECT user_id, login FROM users WHERE user_id = $1`

	addAccountQuery = `INSERT INTO accounts (user_id, title, login, password)
		VALUES ($1, $2, $3, $4)
		RETURNING title
		`
	addTextQuery = `INSERT INTO texts (user_id, title, data)
		VALUES ($1, $2, $3)`

	addBinaryQuery = `INSERT INTO binaries (user_id, title, data)
		VALUES ($1, $2, $3)`

	addCardQuery = `INSERT INTO cards (user_id, title, name, card_number, date_exp, cvc_code)
		VALUES ($1, $2, $3, $4, $5, $6)`

	getByTitleAccountsQuery = `SELECT login, password FROM accounts WHERE user_id= $1 and title = $2`
	getByTitleTextQuery     = `SELECT data FROM texts WHERE user_id= $1 and title = $2`
	getByTitleCardQuery     = `SELECT name, card_number, date_exp, cvc_code FROM cards WHERE user_id= $1 and title = $2`
	getByTitleBinaryQuery   = `SELECT data FROM binaries WHERE user_id= $1 and title = $2`

	getByFullListAccountsQuery = `SELECT title,login,password FROM accounts WHERE user_id= $1`
	getByFullListTextQuery     = `SELECT title,data FROM texts WHERE user_id= $1`
	getByFullListCardQuery     = `SELECT title,card_number,name,date_exp,cvc_code FROM cards WHERE user_id= $1`
	getByFullListBinaryQuery   = `SELECT title,data FROM binaries WHERE user_id= $1`
)
