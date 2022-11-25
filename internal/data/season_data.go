package data

import (
	"encoding/json"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

type SeasonData struct {
	StartDate        time.Time       `json:"start_date"`
}

func ReadSeasonData(path string) (data SeasonData) {
	rd := SeasonData{}
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warnf("error reading season data from file: %v", err)
		return rd
	}

	err = json.Unmarshal([]byte(d), &rd)
	if err != nil {
		log.Warnf("error unmarshaling json data: %v", err)
		return rd
	}
	log.Infof("read season data: %v", rd)

	return rd
}
