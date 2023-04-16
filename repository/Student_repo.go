package repository

import (
	"app-mahasiswa-api2/model"
	"database/sql"
	"fmt"
	"log"
)

type StudentRepo interface {
	GetAll() any
	GetById(id int) any
	Create(newStudent *model.Student) string
	CreateUser(newUser *model.Student) string
	Update(student *model.Student) string
	Delete(id int) string
	DeleteUser(username string) string
}

type studentRepo struct {
	db *sql.DB
}

func (r *studentRepo) GetAll() any {
	var students []model.Student

	query := "select id, name, age, major from students;"
	rows, err := r.db.Query(query)
	fmt.Println(rows)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var student model.Student

		if err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Major); err != nil {
			log.Println(err)
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(students)
	if len(students) == 0 {
		return "no data"
	}

	return students

}

func (r *studentRepo) GetById(id int) any {
	var student model.Student

	query := "SELECT id, name,age,major FROM students WHERE id = $1"

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&student.ID, &student.Name, &student.Age, &student.Major); err != nil {
		log.Println(err)
	}

	if student.ID == 0 {
		return "student not found"
	}

	return student

}

func (r *studentRepo) Create(newStudent *model.Student) string {
	query := "INSERT INTO students ( name,age,major,username) VALUES($1,$2,$3,$4)"
	_, err := r.db.Exec(query, newStudent.Name, newStudent.Age, newStudent.Major, newStudent.UserName)

	if err != nil {
		log.Println(err)
		return "failed to create student"
	}

	// query2:= "INSERT INTO users (id,username,password) VALUES($1,$2,$3)"
	// _,err:= r.db.Exec(query2,newStudent.ID,newStudent.UserName,newStudent.Password)

	return "Student created successfully"
}

func (r *studentRepo) CreateUser(newUser *model.Student) string {

	query := "INSERT INTO users ( username,password) VALUES($1,$2)"
	_, err := r.db.Exec(query, newUser.UserName, newUser.Password)

	if err != nil {
		log.Println(err)
		return " and failed to create user"
	}

	return " and user created successfully"

}

func (r *studentRepo) Update(student *model.Student) string {
	res := r.GetById(student.ID)
	fmt.Println(student)
	if res == "student not found" {
		return res.(string)
	}

	query := "UPDATE students SET name = $1, age = $2, major = $3 WHERE id = $4 ;"
	_, err := r.db.Exec(query, student.Name, student.Age, student.Major, student.ID)

	if err != nil {
		log.Println(err)
		return "failed to update student"
	}

	return fmt.Sprintf("student with id %d updated successfully", student.ID)

}

func (r *studentRepo) Delete(id int) string {
	res := r.GetById(id)
	if res == "student not found" {
		return res.(string)
	}

	query := "DELETE FROM students WHERE id =$1"
	_, err := r.db.Exec(query, id)

	if err != nil {
		log.Println(err)
		return "failed to delete student"
	}

	return fmt.Sprintf("student with id %d deleted successfully", id)
}

// fungsi Delete biasa untuk user
func (r *studentRepo) DeleteUser(username string) string {
	query := "DELETE FROM users WHERE username =$1"
	_, err := r.db.Exec(query, username)
	if err != nil {
		log.Println(err)
		return "failed to delete user"
	}
	return fmt.Sprintf("student with id %s deleted successfully", username)
}

func NewStudentRepo(db *sql.DB) StudentRepo {
	repo := new(studentRepo)

	repo.db = db

	return repo
}
