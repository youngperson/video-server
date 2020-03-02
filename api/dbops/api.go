package dbops

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/video-server/api/defs"
	"github.com/video-server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? and pwd=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // M D y, HH:MM:SS
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info(id, author_id, name, display_ctime) 
	VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	fmt.Println(vid, aid, name, ctime)

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT name, author_id, display_ctime FROM video_info WHERE id=?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&name, &aid, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments(id, video_id, author_id, content) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 时间段内的对某个视频的评论
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	smtmOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments
	INNER JOIN users where comments.author_id = users.id
	AND video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment
	rows, err := smtmOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer smtmOut.Close()
	return res, nil
}
