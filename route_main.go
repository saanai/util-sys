package main

import (
	"net/http"

	"github.com/saanai/util-sys/authentication"
	"github.com/saanai/util-sys/infra/mysql"
	"github.com/saanai/util-sys/util"
	log "github.com/sirupsen/logrus"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := authentication.LoginOrNot(w, r)
	if err != nil {
		util.GenerateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		util.GenerateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	entities, err := mysql.Rc.GetAllThreads()
	data := util.MapEntitiesToDataThreads(entities)
	if err != nil {
		log.Errorf("failed to get all threads. error: %v", err.Error())
		return
	}

	// エラー処理を真面目にしていないが異常時はfalseになるはず
	valid, _ := authentication.LoginOrNot(w, r)
	if valid {
		util.GenerateHTML(w, data, "layout", "private.navbar", "index")
	} else {
		util.GenerateHTML(w, data, "layout", "public.navbar", "index")
	}
}
