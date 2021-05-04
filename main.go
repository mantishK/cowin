package main

import (
	"flag"
	"fmt"

	"github.com/mantishK/cowin/cowin"
)

func main() {
	hubliDistrictID := 278
	blrDistrictID := 265
	puneDistrictID := 363

	var districtFlag = flag.String("d", "blr", "district. blr, hbl, pn")
	var minAgeFlag = flag.Int("a", 18, "min age. 18, 45")
	flag.Parse()

	districtID := hubliDistrictID

	switch *districtFlag {
	case "blr":
		districtID = blrDistrictID
	case "hbl":
		districtID = hubliDistrictID
	case "pn":
		districtID = puneDistrictID
	default:
		panic("incorrect district. Please choose blr, hbl, pn")
	}

	centers := cowin.GetSchedule(districtID, *minAgeFlag)
	if len(centers) == 0 {
		fmt.Println("Sorry, no slots available")
		return
	}
	for _, center := range centers {
		fmt.Println("Name:", center.Name)
		fmt.Println("Address:", center.Address)
		fmt.Println("Block:", center.BlockName)
		fmt.Println("Date:", center.Date)
		fmt.Println("Vaccine:", center.Vaccine)
		fmt.Println("Slots:", center.Slots)
		fmt.Println("=========================================")
		fmt.Print("\n\n")
	}
}
