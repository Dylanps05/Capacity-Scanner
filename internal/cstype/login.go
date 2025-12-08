package cstype

type Password struct {
	Version int    `db:"password_version"`
	Salt    string `db:"password_salt"`
	Hash    string `db:"password_hash"`
}

type Login struct {
	UUID     UserID `db:"id"`
	Username string `db:"username"`
	Password        // Declared anonymously due to pgx v5 limitation
}
