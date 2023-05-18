package postgres

import (
	"auth/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type User struct {
	postgresDB *sqlx.DB
	redisDB    *redis.Client
}

func NewUser(db *sqlx.DB, redis *redis.Client) *User {
	return &User{
		postgresDB: db,
		redisDB:    redis,
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
			return 0, model.ErrorAlreadyExist
		}
		if err.(*pq.Error).Constraint == "users_phone_key" {
			return 0, model.ErrorPhoneAlreadyLinked
		}
		return 0, err
	}

	return userId, nil
}
