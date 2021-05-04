package cowin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Center struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BlockName string    `json:"block_name"`
	Sessions  []Session `json:"sessions"`
}

type Session struct {
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"available_capacity"`
	MinAge            int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

type CenterDetails struct {
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	BlockName         string   `json:"block_name"`
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"available_capacity"`
	MinAge            int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

func GetSchedule(districtID int, minAge int) []CenterDetails {
	date := time.Now().Format("02-01-2006")
	resp, err := http.Get(fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=%d&date=%s", districtID, date))
	if err != nil {
		log.Println(err)
		panic("Failed to fetch")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic("Failed to read response")
	}
	c := struct {
		Center []Center `json:"centers"`
	}{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Println(err)
		panic("Failed to parse json")
	}
	availableCenters := make([]CenterDetails, 0)
	for _, center := range c.Center {
		for _, session := range center.Sessions {
			if session.AvailableCapacity > 0 && minAge >= session.MinAge {
				centerDetails := CenterDetails{
					Name:              center.Name,
					Address:           center.Address,
					BlockName:         center.BlockName,
					AvailableCapacity: session.AvailableCapacity,
					Date:              session.Date,
					MinAge:            session.MinAge,
					Vaccine:           session.Vaccine,
					Slots:             session.Slots,
				}
				availableCenters = append(availableCenters, centerDetails)
			}
		}
	}
	return availableCenters
}
