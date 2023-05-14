package repository

const (
	createUserQuery = `INSERT INTO users (login, password) 
		VALUES ($1, $2) 
		RETURNING user_id, login, password`

	findByLoginQuery = `SELECT user_id, login, password FROM users WHERE login = $1`

	findByIDQuery = `SELECT user_id, login FROM users WHERE user_id = $1`

	addAccountQuery = `INSERT INTO accounts (user_id, title, data)
		VALUES ($1, $2, $3)`

	addTextQuery = `INSERT INTO texts (user_id, title, data)
		VALUES ($1, $2, $3)`

	addBinaryQuery = `INSERT INTO binaries (user_id, title, data)
		VALUES ($1, $2, $3)`

	addCardQuery = `INSERT INTO cards (user_id, title, data)
		VALUES ($1, $2, $3)`

	getByTitleQuery = `SELECT data FROM accounts
		JOIN texts on accounts.user_id = texts.user_id
		JOIN binaries on accounts.user_id = binaries.user_id
		JOIN cards on accounts.user_id = cards.user_id
		WHERE user_id = $1 and title = $2`

	getFullListQuery = `SELECT title, data FROM accounts
		JOIN texts on accounts.user_id = texts.user_id
		JOIN binaries on accounts.user_id = binaries.user_id
		JOIN cards on accounts.user_id = cards.user_id
		WHERE user_id = $1`

	getAllTitlesQuery = `SELECT title FROM accounts
		JOIN texts on accounts.user_id = texts.user_id
		JOIN binaries on accounts.user_id = binaries.user_id
		JOIN cards on accounts.user_id = cards.user_id
		WHERE user_id = $1`
)
