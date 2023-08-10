package main

func main(){
	var p1 Person

	p1.Id=10
	p1.Name="Victoria"
	p1.Email="Victoria@VictoriaSecrets.com"
	p1.Mobile="+91234432211"

	p2:=Person(Id;102,Name:"VS",Email:"vs@vs.com",Mobile:"")


}


type Person struct {
	Id int
	Name string
	Email string
	Mobile string
}