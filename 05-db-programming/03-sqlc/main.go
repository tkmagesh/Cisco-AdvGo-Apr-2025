package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "root:rootuser@tcp(localhost:3306)/go_db_demo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	db := New(conn)

	// inserting the record
	//initializing record to be inserted
	fmt.Println("inserting the record")
	fmt.Scanln()
	newSt := addStudentParams{
		Fname:       "Suresh",
		Lname:       "Kannan",
		DateOfBirth: time.Date(1994, time.August, 14, 23, 51, 42, 0, time.UTC),
		Email:       "suresh@senate.gov",
		Gender:      "Male",
		Address:     "39 Kipling Pass",
	}

	ctx := context.Background()

	sID, err := db.addStudent(ctx, newSt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("addSudent id: %v \n", sID)

	//retreive record by id
	fmt.Println("retreive record by id")
	fmt.Scanln()
	st, err := db.studentByID(ctx, sID)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("studentByID record: %v \n", st)

	//fetching multiple records
	fmt.Println("fetching multiple records")
	fmt.Scanln()
	students, err := db.fetchStudents(ctx)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("fetchStudents count: %v \n", len(students))
	for _, stud := range students {
		fmt.Printf("%+v\n", stud)
	}
}
