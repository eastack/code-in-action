package data

import "time"

type User struct {
	Id       int
	Uuid     string
	Name     string
	Email    string
	Password string
	CreateAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    string
	CreatedAt time.Time
}

// 检查会话有效性
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}
