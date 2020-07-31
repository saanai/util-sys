package data

import (
	"time"

	"github.com/saanai/util-sys/entity"

	"github.com/saanai/util-sys/infra/mysql"
	"github.com/sirupsen/logrus"
)

/*
循環参照を避けるためにentityと別にdataを定義する
*/
type Thread struct {
	Id        int64
	Uuid      string
	Topic     string
	UserId    int64
	CreatedAt time.Time
}

type Post struct {
	Id        int64
	Uuid      string
	Body      string
	UserId    int64
	ThreadId  int64
	CreatedAt time.Time
}

func (thread Thread) NewDataThread(id int64, uuid, topic string, userId int64, createdAt time.Time) *Thread {
	return &Thread{
		Id:        id,
		Uuid:      uuid,
		Topic:     topic,
		UserId:    userId,
		CreatedAt: createdAt,
	}
}

// format the CreatedAt date to display nicely on the screen
func (thread Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// format the CreatedAt date to display nicely on the screen
func (thread Thread) NumReplies() int64 {
	count, err := mysql.Rc.GetNumRepliesOfAnyThread(thread.Id)
	if err != nil {
		logrus.Errorf("db error. error: %v", err.Error())
	}
	return count
}

// Get the user who started this thread
func (thread Thread) User() (user *entity.User) {
	user, err := mysql.Rc.GetUserByUserId(thread.UserId)
	if err != nil {
		logrus.Errorf("db error. error: %v", err.Error())
	}
	return user
}

func (post Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

/*


// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}



// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}



// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}



// Get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}*/
