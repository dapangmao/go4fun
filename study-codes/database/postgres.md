
- using lib/pq directly
```go
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


type result struct {
	title, description string
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=12345 dbname=test sslmode=disable")
	checkErr(err)
	rows, err := db.Query("SELECT title, description FROM film")
	checkErr(err)
	for rows.Next() {
		var r result
		rows.Scan(&r.title, &r.description)
		fmt.Println(r)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
```


- using GORM
```go
package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=12345 dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	db.DropTable(&User{})
	db.CreateTable(&User{})
	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	db.Debug().Save(&User{
		Username:  "adent",
		FirstName: "Arthur",
		LastName:  "Dent",
		Calendar: Calendar{
			Name: "Improbable Events",
			Appointments: []Appointment{
				{Subject: "Spontaneous Whale Generation"},
				{Subject: "Saved from the Vaccuum of Space"},
			},
		},
	})

}

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Calendar  Calendar
}

type Calendar struct {
	gorm.Model
	Name         string
	UserID       uint
	Appointments []Appointment
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
}
```