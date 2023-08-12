package main

func task14() {

}

type Company struct {
	Id        int
	name      string
	website   string
	telephone string
	fax       string
	address   string
}

type Address struct {
	city    string
	line1   string
	street  string
	state   string
	country string
	pincode int
	Map
}

type Map struct {
}
