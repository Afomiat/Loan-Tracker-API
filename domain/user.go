package domain

type User struct {
    ID           string `bson:"_id,omitempty"`
    Email        string `bson:"email"`
    PasswordHash string `bson:"password_hash"`
    IsVerified   bool   `bson:"is_verified"`
}

type VerificationToken struct {
    Token string `bson:"token"`
    Email string `bson:"email"`
}
