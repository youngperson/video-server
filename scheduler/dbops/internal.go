package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmOut, err := dbConn.Prepare("SELECT video_id FROM video_del_rec limit ?")

	var ids []string

	if err != nil {
		return ids, err
	}

	rows, err := stmOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletionRecord error:%v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	defer stmOut.Close()
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmDel, err := dbConn.Prepare("DELETE FROM video_del_rec where video_id=?")
	if err != nil {
		return err
	}

	_, err = stmDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecord error: %v", err)
		return err
	}

	defer stmDel.Close()
	return nil
}
