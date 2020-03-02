package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_infos")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("del", testDeleteUser)
	t.Run("reget", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("tony1", "111")
	if err != nil {
		t.Errorf("Erroe of AddUserCredential: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("tony1")
	if pwd != "111" || err != nil {
		t.Errorf("Erroe of GetUserCredential: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("tony1", "111")
	if err != nil {
		t.Errorf("Erroe of testDeleteUser: %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("tony1")
	if err != nil {
		t.Errorf("Erroe of testReGetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleteing user failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideo)
	t.Run("GetVideo", testGetVideo)
	t.Run("DelVideo", testDeleteVideo)
	t.Run("RegetVideo", testReGetVideo)
}

func testAddVideo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Erroe of AddNewVideo: %v", err)
		return
	}
	tempvid = vi.Id
}

func testGetVideo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Erroe of GetVideoInfo: %v", err)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Erroe of DeleteVideoInfo: %v", err)
	}
}

func testReGetVideo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Erroe of testReGetVideo: %v", err)
	}
}

func TestCommentskFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Erroe of AddNewComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Erroe of testListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Sprintf("coment:%d %v \n", i, ele)
	}
}
