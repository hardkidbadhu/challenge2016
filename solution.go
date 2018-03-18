package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

//Distributor struct holds the distributor records
type Distributor struct {
	Name            string
	SDis            []string //sub distributor associated with this distributor
	AccessAreas     []string
	RestrictedAreas []string
}

//Area structure to hold the area data from csv
type Area struct {
	Country  string
	Province string
	City     string
}

type Err struct {
	Name  string
}

// Error function for satisfying error interface
func (err Err) Error() string {
	return fmt.Sprintf("Message: %s", err.Name)
}

//map to hold records for processing
var CountryMap, ProvinceMap, CityMap map[string]*Area

//Distributor map for holding distributor map

var DistributorMap map[string]*Distributor

//structure for holding city data
type City struct {
}

func main() {

	log.Println("Distributor Solution....")

	log.Println("Processing the csv file...")

	CountryMap, ProvinceMap, CityMap = ProcessFile("cities.csv")

	log.Println("File processed....")

	log.Println("Its time to configure distributor!!")

	//var to hold no.of distributors
	var dis int

	log.Println("Select the no.of distributors to configure")

	fmt.Scanf("%d", &dis)

	//function to configure distributor
	err, disMap := ConfigureDistributors(dis)
	if err != nil {
		log.Println("Configure the Distributor again!..")
		err, disMap = ConfigureDistributors(dis)
		if err != nil {
			log.Println("Sorry try after some time.....")
			return
		}
	}

}

func ProcessFile(fileName string) (map[string]*Area, map[string]*Area, map[string]*Area) {

	//open a file
	citiesFile, _ := os.Open(fileName)
	defer citiesFile.Close()

	readerObj := csv.NewReader(bufio.NewReader(citiesFile))

	//areas hold the data in structure
	areas := []Area{}

	for {

		readLine, err := readerObj.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if readLine[0] == "City Code" {
			continue
		}

		area := Area{}
		area.City = readLine[0]
		area.Province = readLine[1]
		area.Country = readLine[2]
		areas = append(areas, area)

	}

	//putting all the areas data into map
	//specific to country province and city
	CountryMap = make(map[string]*Area)
	ProvinceMap = make(map[string]*Area)
	CityMap = make(map[string]*Area)

	for i := range areas {
		CountryMap[areas[i].Country] = &areas[i]
	}

	for j := range areas {
		ProvinceMap[areas[j].Province] = &areas[j]
	}

	for k := range areas {
		CityMap[areas[k].City] = &areas[k]
	}

	return CountryMap, ProvinceMap, CityMap
}

func ConfigureDistributors(noOfDis int) (error, map[string]*Distributor) {

	var Distributors []*Distributor
	var err Err
	for i := 1; i == noOfDis; i++ {
		log.Printf("Configure the Distributor %d profile", i)

		log.Println("Enter the name of the distributor.....")
		dis := Distributor{}
		fmt.Scanf("%s", dis.Name)

		log.Println("Enter the no of access areas you want to grant for the distributor....")
		access := 0
		fmt.Scanf("%d", &access)
		//adding the access areas
		for j := 1; j == access; j++ {
			var level int
			var accessArea string
			log.Printf("Add the access area %d", j)
			log.Printf("INFO: Give us the level \n 0 - Country level \n 1 - Province level \n 2 - City level")
			fmt.Scanf("%d", &level)
			switch level {
			case 0:
				log.Println("Level 0 you must enter the county code")
				fmt.Scanf("%s", &accessArea)
				if _, ok := CountryMap[accessArea]; !ok {
					err.Name = "Invalid Country code"
					return err, nil
				}
				dis.AccessAreas = append(dis.AccessAreas, accessArea)
			case 1:
				log.Println("Level 1 you must enter the Province code")
				fmt.Scanf("%s", &accessArea)
				if _, ok := ProvinceMap[accessArea]; !ok {
					err.Name = "Invalid Province code"
					return err, nil
				}
				dis.AccessAreas = append(dis.AccessAreas, accessArea)
			case 2:
				log.Println("Level 2 you must enter the city code")
				fmt.Scanf("%s", &accessArea)
				if _, ok := CityMap[accessArea]; !ok {
					err.Name = "Invalid city code"
					return err, nil
				}
				dis.AccessAreas = append(dis.AccessAreas, accessArea)
			}

		}

	}
}

