package data

import (
	"encoding/json"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

type RotationEntry struct {
	Name     string `json:"name"`
	Location int    `json:"location"`
	Comments string `json:"comments"`
}

type RotationData struct {
	StartDate        time.Time       `json:"start_date"`
	RotationComplete bool            `json:"rotation_complete"`
	GearList         []string        `json:"gear_list"`
	LocationList     []string        `json:"location_list"`
	ContentRotation  []RotationEntry `json:"content_rotation"`
	GearRotation     []int           `json:"gear_rotation"`
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
