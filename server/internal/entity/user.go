package entity

type User struct {
	ID       string
	Login    string
	Password string
}

// LoginParole
type LoginParole struct {
	ID       string
	Login    string
	Password string
}

// Text
type Text struct {
	TextData string
}

// Binary
type Binary struct {
	BinaryData byte
}

// Card
type Card struct {
	CardNumber     string
	ExpirationData string
	CardHolderName string
	CVCCode        string
}
