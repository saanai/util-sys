package authentication

import (
	"github.com/saanai/util-sys/entity"
	"github.com/saanai/util-sys/infra/mysql"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostFormValue("email")
	user, err := mysql.Rc.GetUserByEmail(email)
	if err != nil {
		log.Errorf("failed to get user by email. error: %v", err)
	}

	password := r.PostFormValue("password")
	//hash, err := bcrypt.GenerateFromPassword([]byte(requestPassword), bcrypt.DefaultCost)
	//if err != nil {
	//	log.Errorf("hash password generation failed. error: %v", err)
	//}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// failed to login
		http.Redirect(w, r, "/login", 302)
	} else {
		// success to login
		session := &entity.Session{
			Uuid: user.Uuid,
			Email: user.Email,
			UserId: user.Id,
			CreatedAt: time.Now(),
		}
		err = mysql.Rc.CreateSession(session)
		if err != nil {
			log.Errorf("failed to create session. error: %v", err)
		}
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	}

}
