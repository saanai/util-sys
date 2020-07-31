package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/saanai/util-sys/entity"
	"github.com/sirupsen/logrus"
)

type client struct {
	db *sql.DB
}

var c *client

func init() {
	database, err := sql.Open("mysql", "app_be:hogehoge@/utilsys?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		logrus.Error("mysql client initialization failed")
	}
	c = &client{
		db: database,
	}
}

// Create a new session for an existing user
func (client *client) createUser(user *entity.User) (err error) {
	statement := "INSERT INTO users (uuid, email, name, password, created_at) VALUES (?,?,?,?,?)"
	stmt, err := client.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&user.Uuid, &user.Email, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return err
	}
	return err
}

// Get a single user given the email
func (client *client) getUserByEmail(email string) (user *entity.User, err error) {
	user = &entity.User{}
	statement := "SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?"
	err = client.db.QueryRow(statement, email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func (client *client) createSession(session *entity.Session) (err error) {
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?)"
	stmt, err := client.db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return err
	}
	return err
}

// Get a single user given the email
func (client *client) checkSession(uuid string) (valid bool, err error) {
	session := &entity.Session{}
	statement := "SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?"
	err = client.db.QueryRow(statement, uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return false, err
	}
	return true, err
}

// Get all threads in the database and returns it
func (client *client) getAllThreads() (threads []entity.Thread, err error) {
	statement := "SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC"
	rows, err := client.db.Query(statement)
	if err != nil {
		return
	}
	for rows.Next() {
		thread := &entity.Thread{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt); err != nil {
			return
		}
		threads = append(threads, *thread)
	}
	rows.Close()
	return threads, err
}

// get the number of posts in a thread
func (client *client) getNumRepliesOfAnyThread(threadId int64) (count int64, err error) {
	statement := "SELECT count(*) FROM posts where thread_id = ?"
	err = client.db.QueryRow(statement, threadId).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, err
}

func (client *client) getUserByUserId(userId int64) (user *entity.User, err error) {
	user = &entity.User{}
	statement := "SELECT id, uuid, name, email, created_at FROM users WHERE id = ?"
	err = client.db.QueryRow(statement, userId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, err
}
