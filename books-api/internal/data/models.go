package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

const dbTimeout = 3 * time.Second

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User:  User{},
		Token: Token{},
	}
}

type Models struct {
	User  User
	Token Token
}
type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     Token     `json:"token"`
}

func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT id, email, password, first_name, last_name, created_at, updated_at FROM users order by last_name"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var user User
	query := "select * from users where email = $1"
	row := db.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *User) GetById(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var user User
	query := "select * from users where id = $1"
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := "update users set email = $1, first_name = $2, last_name = $3, updated_at = $4 where id = $5"
	_, err := db.ExecContext(ctx, stmt, u.Email, u.FirstName, u.LastName, u.UpdatedAt, u.Id)
	if err != nil {
		return err
	}
	return nil
}
func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := "delete from users where id = $1"
	_, err := db.ExecContext(ctx, stmt, u.Id)
	if err != nil {
		return err
	}
	return nil
}
func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}
	var id int
	stmt := "insert into users (email, password, first_name, last_name, created_at, updated_at) values ($1, $2, $3, $4, $5, $6) returning id"
	err = db.QueryRowContext(ctx, stmt, user.Email, hashedPassword, user.FirstName, user.LastName, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := "update users set password = $1, updated_at = $2 where id = $3"
	_, err = db.ExecContext(ctx, stmt, hashedPassword, time.Now(), u.Id)
	if err != nil {
		return err
	}
	return nil
}
func (u *User) VerifyPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

type Token struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"token_hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Expiry    time.Time `json:"expiry"`
}

func (t *Token) GetByToken(plainTextToken string) (*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var token Token
	query := "select * from tokens where token = $1"
	row := db.QueryRowContext(ctx, query, plainTextToken)
	err := row.Scan(&token.Id, &token.UserId, &token.Email, &token.Token, &token.TokenHash, &token.CreatedAt, &token.UpdatedAt, &token.Expiry)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
func (t *Token) GetUserByToken(token Token) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := "select * from users where id = $1"
	var user User
	row := db.QueryRowContext(ctx, query, token.UserId)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (t *Token) GenerateToken(userId int, ttl time.Duration) (*Token, error) {

	token := &Token{
		UserId: userId,
		Expiry: time.Now().Add(ttl),
	}

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return nil, err

	}
	token.Token = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Token))
	token.TokenHash = hash[:]

	return token, nil
}
func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {
	AuthHeader := r.Header.Get("Authorization")
	if AuthHeader == "" {
		return nil, errors.New("Authorization header is empty")
	}

	headerParts := strings.Split(AuthHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("Authorization header is invalid")
	}

	tokenString := headerParts[1]
	if len(tokenString) != 26 {
		return nil, errors.New("wrong token length")
	}

	token, err := t.GetByToken(tokenString)
	if err != nil {
		return nil, errors.New("No token found")
	}

	if token.Expiry.Before(time.Now()) {
		return nil, errors.New("Token is expired")
	}

	user, err := token.GetUserByToken(*token)
	if err != nil {
		return nil, errors.New("No matching user found")
	}

	return user, nil
}
func (t *Token) Insert(token Token, user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := "delete from tokens where user_id = $1"
	_, err := db.ExecContext(ctx, stmt, token.UserId)
	if err != nil {
		return err
	}
	token.Email = user.Email
	stmt = "insert into tokens (user_id, email, token, token_hash, created_at, updated_at, expiry) values ($1, $2, $3, $4, $5, $6, $7) returning id"
	_, err = db.ExecContext(ctx, stmt, token.UserId, token.Email, token.Token, token.TokenHash, time.Now(), time.Now(), token.Expiry)
	if err != nil {
		return err
	}
	return nil
}
func (t *Token) DeleteToken(token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := "delete from tokens where token = $1"
	_, err := db.ExecContext(ctx, stmt, token)
	if err != nil {
		return err
	}
	return nil
}
func (t *Token) IsValidToken(tokenString string) (bool, error) {
	token, err := t.GetByToken(tokenString)
	if err != nil {
		return false, errors.New("No token found")
	}
	_, err = t.GetUserByToken(*token)
	if err != nil {
		return false, errors.New("No matching user found")
	}
	if token.Expiry.Before(time.Now()) {
		return false, errors.New("Token is expired")
	}
	return true, nil
}
