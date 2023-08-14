package main

import "fmt"

func task14() {

	//check if promioted field address can be direcly accesible
	var company Company = Company{
		Id:        1,
		name:      "Victoria Secrets & Co",
		website:   "https://www.victoriassecret.com/in/",
		telephone: "67686534553",
		fax:       "738393393933",
		Address: Address{ //when using promoted field we gotta use the struct name
			city:    "Bangalore",
			line1:   "manyata tech park",
			street:  "embassy road",
			state:   "Karnataka",
			country: "India",
			pincode: 56008,
			Map: Map{ //when using promoted field we gotta use the struct name
				lat: 45.56,
				lan: 35.67,
			},
		},
	}

	fmt.Println("Address is directly accesible")
	fmt.Println(company.city)
	fmt.Println("Map is also directly accesible so promoted fields of promoted fields and so on are directly accesible")
	fmt.Println(company.lan)

}

type Company struct {
	Id        int
	name      string
	website   string
	telephone string
	fax       string
	Address   //promoted field
}

type Address struct {
	city    string
	line1   string
	street  string
	state   string
	country string
	pincode int
	Map     //promoted field
}

type Map struct {
	lat, lan float64
}
