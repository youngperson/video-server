package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/video-server/api/defs"
)

func InserSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err := stmIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	defer stmIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")
	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string
	stmOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var ttlint int64
	if res, ttlint := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}

	rows, err := stmOut.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Panicf("retrive sessions error:%s", err)
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
		} else {
			return nil, err
		}
	}
	defer stmOut.Close()
	return m, nil
}

func DeleteSession(sid string) error {
	stmOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil {
		return err
	}

	if _, err := stmOut.Query(sid); err != nil {
		return err
	}
	defer stmOut.Close()
	return nil
}
