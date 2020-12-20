package systemLog

import (
	"fmt"
	"log-collector/data"
	"log-collector/model/_util/pageInfo"
)

func GetByPageInfo(pi *pageInfo.Request) ([]SystemLog, *pageInfo.Response, error) {
	return nil, nil, nil
}

func Create(s *SystemLog) error {
	conn := *data.GetQuestDBConn()
	_, err := conn.Write(s.ToInfluxSQL())
	if err != nil {
		return fmt.Errorf("conn.Write(s.ToInfluxSQL()): %s", err)
	}

	return nil
}
