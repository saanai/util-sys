package authentication

import (
	"github.com/saanai/util-sys/infra/mysql"
	"net/http"
)

// errとvalidの処理の棲み分けを真面目に考えていないので要検討（エラー処理が杜撰）
func LoginOrNot(w http.ResponseWriter, r *http.Request) (valid bool, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		return false, err
	}
	valid, err = mysql.Rc.CheckSession(cookie.Value)
	return valid, err
}
