package users

import (
	"errors"
	"gzfs/Go~Edita/internal/database"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
		name VARCHAR(255), 
		email VARCHAR(255) UNIQUE
		)`

	_, err := tursoDB.TursoDB.Exec(createUsersTableQuery)
	return err
}

func (userStorage *UserStorage) Create(createUser *User) error {
	if createUser.Email == "" {
		return errors.New("email is required")
	}

	if createUser.Name == "" {
		return errors.New("name is required")
	}

	createUserQuery := "INSERT INTO users (id, name, email) VALUES (?, ?, ?)"
	_, err := userStorage.userDB.TursoDB.Exec(createUserQuery, createUser.ID, createUser.Name, createUser.Email)
	return err
}

func (userStorage *UserStorage) GetByEmail(userEmail string) (*User, error) {
	var getUser User
	getUserQuery := "SELECT id, name, email FROM users WHERE email = ?"
	err := userStorage.userDB.TursoDB.QueryRow(getUserQuery, userEmail).Scan(&getUser.ID, &getUser.Name, &getUser.Email)
	return &getUser, err
}

func (userStorage *UserStorage) UpdateByEmail(userEmail string, updateUser *User) error {
	updateUserQuery := "UPDATE users SET name = ?, email = ? WHERE email = ?"
	_, err := userStorage.userDB.TursoDB.Exec(updateUserQuery, updateUser.Name, updateUser.Email, userEmail)
	return err
}

func (userStorage *UserStorage) DeleteByEmail(userEmail string) error {
	deleteUserQuery := "DELETE FROM users WHERE email = ?"
	_, err := userStorage.userDB.TursoDB.Exec(deleteUserQuery, userEmail)
	return err
}
