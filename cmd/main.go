package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/mantishK/cowin/cowin"
	"github.com/mantishK/cowin/user"
)

func main() {
	var districtFlag = flag.String("d", "blr", "district. blr, hbl, pn")
	var minAgeFlag = flag.Int("a", 18, "min age. 18, 45")
	var operationType = flag.String("operation", "display", "openration type. display or export")

	var users user.UserFlags
	flag.Var(&users, "usr", "User details. Should be in the form - abc@gmail.com:blr:45")
	flag.Parse()

	if *operationType == "display" {

		districtID := cowin.GetDistrictID(*districtFlag)
		centers := cowin.GetSchedule(districtID, *minAgeFlag)
		if len(centers) == 0 {
			fmt.Println("Sorry, no slots available")
			return
		}
		details := cowin.GetFormattedCenters(centers)
		fmt.Println(details)
	} else if *operationType == "export" {
		for _, usr := range users {

			userDetails := strings.Split(usr, ":")
			email := userDetails[0]
			district := userDetails[1]
			minAge, err := strconv.Atoi(userDetails[2])
			if err != nil {
				panic("min age not number")
			}

			districtID := cowin.GetDistrictID(district)
			centers := cowin.GetSchedule(districtID, minAge)
			if len(centers) != 0 {
				user.SendMail(email, cowin.GetFormattedCenters(centers))
			}
		}
	}
}
