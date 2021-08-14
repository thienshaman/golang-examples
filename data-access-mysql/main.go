package main

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/data-access/entity"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "admin:12345@tcp(127.0.0.1:3306)/jv44")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connnected!")

	var students1, e1 = findAllStudent()
	fmt.Println(e1)
	fmt.Println(students1)

	var students2, e2 = findStudentByRank("C")
	fmt.Println(e2)
	fmt.Println(students2)

	var student, e3 = findStudentById(365179528)
	fmt.Println(e3)
	fmt.Println(student)

	var id, e4 = addStudent(entity.Student{
		Name:  "Do Phu Thien",
		Email: "thienshaman@gmail.com",
		Rank:  "A",
	})
	fmt.Println(e4)
	fmt.Println(id)

}

func findAllStudent() ([]entity.Student, error) {
	var students []entity.Student
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		return nil, fmt.Errorf("findAllStudent: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var student entity.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Rank); err != nil {
			return nil, fmt.Errorf("findAllStudent: %v", err)
		}
		students = append(students, student)
	}

	return students, nil
}

func findStudentByRank(rank string) ([]entity.Student, error) {
	var students []entity.Student
	rows, err := db.Query("SELECT * FROM student WHERE rank = ?", rank)
	if err != nil {
		return nil, fmt.Errorf("findStudentByRank: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var student entity.Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Rank); err != nil {
			return nil, fmt.Errorf("findStudentByRank: %v", err)
		}
		students = append(students, student)
	}

	return students, nil
}

func findStudentById(id int) (entity.Student, error) {
	var student entity.Student
	row := db.QueryRow("SELECT * FROM student WHERE id = ?", id)

	if err := row.Scan(&student.ID, &student.Name, &student.Email, &student.Rank); err != nil {
		return student, fmt.Errorf("findStudentById: %v", err)
	}
	return student, nil

}

func addStudent(student entity.Student) (int64, error) {
	result, err := db.Exec("INSERT INTO student (name, email, rank) VALUE (?, ?, ?)", student.Name, student.Email, student.Rank)
	if err != nil {
		return 0, fmt.Errorf("addStudent: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addStudent: %v", err)
	}
	return id, nil
}
