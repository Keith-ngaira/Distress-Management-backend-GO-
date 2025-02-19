package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"errors"
)

// NullTime is a wrapper around sql.NullTime that implements JSON marshaling
type NullTime struct {
	sql.NullTime
}

// MarshalJSON implements json.Marshaler
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}

// UnmarshalJSON implements json.Unmarshaler
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t != nil {
		nt.Valid = true
		nt.Time = *t
	} else {
		nt.Valid = false
	}
	return nil
}

// Value implements the driver.Valuer interface
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// Scan implements the sql.Scanner interface
func (nt *NullTime) Scan(value interface{}) error {
	var nt2 sql.NullTime
	if err := nt2.Scan(value); err != nil {
		return err
	}
	nt.Time = nt2.Time
	nt.Valid = nt2.Valid
	return nil
}

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password,omitempty"`
	Role       string    `json:"role"`
	Department string    `json:"department"`
	Active     bool      `json:"active"`
	LastLogin  NullTime  `json:"last_login"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MarshalJSON implements custom JSON marshaling for User
func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID         int64     `json:"id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		Password   string    `json:"password,omitempty"`
		Role       string    `json:"role"`
		Department string    `json:"department"`
		Active     bool      `json:"active"`
		LastLogin  NullTime  `json:"last_login"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		Role:       u.Role,
		Department: u.Department,
		Active:     u.Active,
		LastLogin:  u.LastLogin,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	})
}

func (u *User) Create(db *sql.DB) error {
	log.Printf("DEBUG: Creating user with email: %s", u.Email)
	log.Printf("DEBUG: Password length: %d", len(u.Password))

	if strings.TrimSpace(u.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(u.Password) == "" {
		return errors.New("password is required")
	}

	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Set default values
	u.Active = true
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.LastLogin = NullTime{}

	query := `INSERT INTO users 
		(name, email, password, role, department, active, last_login, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, NULL, ?, ?)`

	result, err := db.Exec(query,
		u.Name,
		u.Email,
		u.Password,
		u.Role,
		u.Department,
		u.Active,
		u.CreatedAt,
		u.UpdatedAt,
	)

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email already exists")
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (u *User) HashPassword() error {
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Store the hashed password
	u.Password = string(hashedPassword)
	return nil
}

func GetUser(db *sql.DB, id int64) (*User, error) {
	u := &User{}

	query := `SELECT id, name, email, role, department, active, last_login, created_at, updated_at 
		FROM users WHERE id = ?`
	err := db.QueryRow(query, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Role,
		&u.Department,
		&u.Active,
		&u.LastLogin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	u := &User{}

	query := `SELECT id, name, email, password, role, department, active, last_login, created_at, updated_at 
		FROM users WHERE email = ?`
	err := db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.Department,
		&u.Active,
		&u.LastLogin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) ComparePassword(password string) error {
	log.Printf("DEBUG: Stored hash: %s", u.Password)
	log.Printf("DEBUG: Input password: %s", password)
	log.Printf("DEBUG: Hash length: %d, Input length: %d", len(u.Password), len(password))

	// Make sure we have a valid bcrypt hash
	if !strings.HasPrefix(u.Password, "$2a$") {
		log.Printf("DEBUG: Invalid hash format - doesn't start with $2a$")
		return errors.New("invalid password format")
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Printf("DEBUG: Password comparison failed: %v", err)
		return err
	}

	log.Printf("DEBUG: Password comparison succeeded")
	return nil
}

func (u *User) Update(db *sql.DB) error {
	query := `UPDATE users SET 
		name = ?, 
		email = ?, 
		role = ?, 
		department = ?, 
		active = ?, 
		last_login = ?, 
		updated_at = ?
		WHERE id = ?`

	u.UpdatedAt = time.Now()
	
	_, err := db.Exec(query,
		u.Name,
		u.Email,
		u.Role,
		u.Department,
		u.Active,
		u.LastLogin,
		u.UpdatedAt,
		u.ID,
	)
	return err
}

func (u *User) UpdatePassword(db *sql.DB, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `UPDATE users SET password = ?, updated_at = NOW() WHERE id = ?`
	_, err = db.Exec(query, string(hashedPassword), u.ID)
	return err
}

func GetUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, name, email, role, department, active, last_login, created_at, updated_at 
		FROM users ORDER BY created_at DESC`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Role,
			&u.Department,
			&u.Active,
			&u.LastLogin,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
