package cowin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var districtIDMap = map[string]int{
	"hbl":  278,
	"blr":  265,
	"pn":   363,
	"bbmp": 294,
}

type Center struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	BlockName string    `json:"block_name"`
	Sessions  []Session `json:"sessions"`
}

type Session struct {
	Date              string   `json:"date"`
	AvailableCapacity float64  `json:"available_capacity"`
	MinAge            int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

type CenterDetails struct {
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	BlockName         string   `json:"block_name"`
	Date              string   `json:"date"`
	AvailableCapacity float64  `json:"available_capacity"`
	MinAge            int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

func GetSchedule(districtID int, minAge int) []CenterDetails {
	date := time.Now().Format("02-01-2006")
	req, err := http.NewRequest("GET", fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id=%d&date=%s", districtID, date), nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	client := http.DefaultClient
	resp, err := client.Do(req)
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

func GetFormattedCenters(centers []CenterDetails) string {
	details := ""
	for _, center := range centers {
		details += fmt.Sprintln("Name:", center.Name)
		details += fmt.Sprintln("Address:", center.Address)
		details += fmt.Sprintln("Block:", center.BlockName)
		details += fmt.Sprintln("Date:", center.Date)
		details += fmt.Sprintln("Vaccine:", center.Vaccine)
		details += fmt.Sprintln("Slots:", center.Slots)
		details += fmt.Sprintln("Min Age:", center.MinAge)
		details += fmt.Sprintln("=========================================")
		details += fmt.Sprint("\n\n")
	}
	return details
}

func GetDistrictID(district string) int {
	return districtIDMap[district]
}
