package postgres

import (
	"auth/internal/model"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type User struct {
	postgresDB *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		postgresDB: db,
	}
}

func (u *User) Create(user model.UserInput) (int, error) {
	var userId int
	query := `INSERT INTO users (name, surname,position,location,birth_date, gender, email, phone, password,role) 
			  VALUES ($1,$2,$3 ,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	err := u.postgresDB.QueryRowx(query,
		user.Name,
		user.Surname,
		user.Position,
		user.Location,
		user.BirthDate,
		user.Gender,
		user.Email,
		user.Phone,
		user.Password,
		user.Role).Scan(&userId)
	if err != nil {
		if err.(*pq.Error).Constraint == "users_email_key" {
			return 0, model.ErrorUserAlreadyExist
		}
		if err.(*pq.Error).Constraint == "users_phone_key" {
			return 0, model.ErrorUserPhoneAlreadyLinked
		}
		return 0, err
	}

	return userId, nil
}

func (u *User) GetByEmail(email string) (model.User, error) {
	query := `SELECT 
    			*
			   FROM users 
			   WHERE email = $1`

	var user model.User
	err := u.postgresDB.Get(&user, query, email)
	//check user is found
	if err == sql.ErrNoRows {
		return model.User{}, model.ErrorUserNotFound
	} else if err != nil {
		return model.User{}, err
	}
	return user, nil
}
