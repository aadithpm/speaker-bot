package data

import (
	"encoding/json"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

type FeaturedContentEntry struct {
	Name          	string 	`json:"name"`
	Location      	string 	`json:"location"`
	MasterAvailable	bool 	`json:"master_available"`
	Comments      	string	`json:"comments"`
}

type FeaturedContentData struct {
	StartDate        time.Time              `json:"start_date"`
	ContentRotation  []FeaturedContentEntry `json:"content_rotation"`
}

func ReadFeaturedContentData(path string) (data FeaturedContentData) {
	rd := FeaturedContentData{}
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Warnf("error reading featured content data from file: %v", err)
		return rd
	}

	err = json.Unmarshal([]byte(d), &rd)
	if err != nil {
		log.Warnf("error unmarshaling json data: %v", err)
		return rd
	}
	log.Infof("read featured content data: %v", rd)

	return rd
}
