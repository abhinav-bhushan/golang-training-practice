package handlers

import (
	"demo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type EmployeeHandler struct {}
var auto *models.AutoIncrement=new (models.AutoIncrement)   //auto incrementing id


func (eh *EmployeeHandler)AddEmployee(w http.ResponseWriter,r *http.Request){
		if r.Method=="POST"{
			e:=new(models.Employee)
			e.Id=auto.Increment()   //automatically storing ID and incrementing 
			err:=json.NewDecoder(r.Body).Decode(e)

			if err!=nil{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			if err:=e.Validate();err!=nil{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return

			}
			e.Status="Active"
			e.LastModified=time.Now().UTC().Unix()

			if bytes,err:=e.ToBytes();err!=nil{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
			} else {
				f,err:=os.OpenFile("employee.txt",os.O_CREATE| os.O_APPEND| os.O_WRONLY,0644)   //FOR APPENDING INTO THE FILE
				if err!=nil{
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				}
				_,err=f.Write(bytes)
				if err!=nil{
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				} else{
					w.Write([]byte("Data succesfully stored in the file"))
					w.WriteHeader(201)
					return
				}
			}



		}
}

func(eh* EmployeeHandler)DeleteEmployee(w http.ResponseWriter,r *http.Request){
	values:=r.URL.Query()
	id:=values.Get("id")

	if id==""{
		w.Write([]byte("invalid id"))
		w.WriteHeader(400)
		return
	}
	fmt.Fprintln(w, "id query param is->",id)


}



