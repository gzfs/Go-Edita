package users

import (
	"gzfs/Go~Edita/internal/database"
	"time"
)

type User struct {
	ID         string    `json:"id" validate:"required"`
	First_Name string    `json:"first_name" validate:"required"`
	Last_Name  string    `json:"last_name" validate:"required"`
	Username   string    `json:"username" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Admin      bool      `json:"admin"`
	Password   string    `json:"password" validate:"required"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type UserStorage struct {
	userDB *database.TursoStorage
}

func NewUserStorage(tursoDB *database.TursoStorage) *UserStorage {
	return &UserStorage{
		tursoDB,
	}
}

func InitUsersTable(tursoDB *database.TursoStorage) error {
	createUsersTableQuery := `CREATE TABLE IF NOT EXISTS Users (
		id VARCHAR(255) PRIMARY KEY, 
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		username VARCHAR(255),
		email VARCHAR(255) UNIQUE,
		password VARCHAR(60),
		is_admin BOOLEAN,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
		)`

	_, err := tursoDB.TursoDB.Exec(createUsersTableQuery)
	return err
}

func (userStorage *UserStorage) Create(createUser *User) error {
	createUserQuery := "INSERT INTO users (id, first_name, last_name, username, email, password, is_admin, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := userStorage.userDB.TursoDB.Exec(createUserQuery, createUser.ID, createUser.First_Name, createUser.Last_Name, createUser.Username, createUser.Email, createUser.Password, createUser.Admin, time.Now(), time.Now())
	return err
}

func (userStorage *UserStorage) GetByEmail(userEmail string) (*User, error) {
	var getUser User
	getUserQuery := "SELECT id, first_name, last_name, username, email, password, is_admin, created_at, updated_at FROM users WHERE email = ?"
	err := userStorage.userDB.TursoDB.QueryRow(getUserQuery, userEmail).Scan(&getUser.ID, &getUser.First_Name, &getUser.Last_Name, &getUser.Username, &getUser.Email, &getUser.Password, &getUser.Admin, &getUser.Created_At, &getUser.Updated_At)
	return &getUser, err
}

func (userStorage *UserStorage) UpdateByEmail(userEmail string, updateUser *User) error {
	updateUserQuery := "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE email = ?"
	_, err := userStorage.userDB.TursoDB.Exec(updateUserQuery, updateUser.First_Name, updateUser.Last_Name, updateUser.Email, userEmail)
	return err
}

func (userStorage *UserStorage) DeleteByEmail(userEmail string) error {
	deleteUserQuery := "DELETE FROM users WHERE email = ?"
	_, err := userStorage.userDB.TursoDB.Exec(deleteUserQuery, userEmail)
	return err
}
