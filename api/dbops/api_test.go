package dbops

import "testing"

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
