package repository

const (
	createUserQuery = `INSERT INTO users (login, password) 
		VALUES ($1, $2) 
		RETURNING user_id, login, password`

	findByLoginQuery = `SELECT user_id, login, password FROM users WHERE login = $1`

	findByIDQuery = `SELECT user_id, login FROM users WHERE user_id = $1`
)
