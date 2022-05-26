package repository

import (
	"fmt"
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	// load data user
	records, err := u.db.Load("users")
	if err != nil {
		// jika error, maka bikin data user baru
		records = [][]string{
			{"username", "password", "loggedin", "role"},
		}
		// simpan data baru ke data
		err = u.db.Save("users", records)
		if err != nil {
			return nil, err
		}
	}

	// jika data user ada,maka load data tersebut
	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		// convert nilai bool
		loggedin, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		// load data user dan append ke variable result
		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: loggedin,
			Role:     records[i][3],
		}

		result = append(result, user)
	}

	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	// panggil function loadOrCreate untuk melakukan load data
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	// ambil data user melalui function selectAll atau loadOrCreate
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	// looping data user dan cocokkan username dan password
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &username, nil
		}
	}

	// jika username dan password tidak cocok, maka return error
	return nil, fmt.Errorf("Login Failed")
}

func (u *UserRepository) Save(users []User) error {
	// sipakan dulu header csv file
	records := [][]string{
		{"username", "password", "loggedin", "role"},
	}

	// loop semua data user yg didapat dari param function
	for i := 0; i < len(users); i++ {
		// append data user ke records
		records = append(records, []string{
			users[i].Username,
			users[i].Password,
			strconv.FormatBool(users[i].Loggedin),
			users[i].Role,
		})
	}

	// simpan data user ke file csv
	return u.db.Save("users", records)

}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	// load data user menggunakan function selectAll
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	// loop data dan cocokan username lalu ambil role
	for _, user := range users {
		if user.Username == username {
			return &user.Role, nil
		}
	}

	// jika username tidak ditemukan, maka return error
	return nil, fmt.Errorf("Faile to get user role")
}
