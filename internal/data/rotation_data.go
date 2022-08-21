package data

import (
	"encoding/json"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

type RotationEntry struct {
	Name     string `json:"name"`
	Gear     int    `json:"gear"`
	Location int    `json:"location"`
}

type RotationData struct {
	StartDate    time.Time       `json:"start_date"`
	RotationType string          `json:"rotation_type"`
	GearList     []string        `json:"gear_list"`
	LocationList []string        `json:"location_list"`
	Rotation     []RotationEntry `json:"rotation"`
}

func ReadRotationData(path string) (data RotationData) {
	rd := RotationData{}
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warnf("error reading rotation data from file: %v", err)
		return rd
	}

	err = json.Unmarshal([]byte(d), &rd)
	if err != nil {
		log.Warnf("error unmarshaling json data: %v", err)
		return rd
	}
	log.Infof("read rotation data: %v", rd)

	return rd
}
