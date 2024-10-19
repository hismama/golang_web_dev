package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = cookieLength
	http.SetCookie(w, c)
	//if the user exists already, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	c.MaxAge = cookieLength
	s := dbSessions[c.Value]
	_, ok := dbUsers[s.un]
	return ok
}

func cleanSessions() {
	fmt.Println("Cleaning sessions...")
	showSessions() // for demo
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (sessionLength) || v.un == "" {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("Cleaning done!")
	showSessions() // for demo
}

func showSessions() {
	fmt.Println("*************")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
