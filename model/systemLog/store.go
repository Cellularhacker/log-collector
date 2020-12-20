package systemLog

import (
	"database/sql"
	"log-collector/model/_util/pageInfo"

	_ "github.com/lib/pq"
)

func GetByPageInfo(pi *pageInfo.Request) ([]SystemLog, *pageInfo.Response, error) {
	return nil, nil, nil
}

func Create(s *SystemLog) error {
	db, err := sql.Open("postgres", "host=localhost port=8812 user=admin password=quest dbname=qdb sslmode=disable")
	if err != nil {
		return err
	}

	defer db.Close()

	rows, err := db.Query(s.ToPQSQL())
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
