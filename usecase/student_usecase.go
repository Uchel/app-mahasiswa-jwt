package usecase

import (
	"app-mahasiswa-api2/model"
	"app-mahasiswa-api2/repository"
)

type StudentUsecase interface {
	FindStudents() any
	FindStudentById(id int) any
	Register(newStudent *model.Student) string
	Edit(student *model.Student) string
	Unreg(id int) string
	UnregUser(username string) string
}

type studentUsecase struct {
	studentRepo repository.StudentRepo
}

func (u *studentUsecase) FindStudents() any {

	return u.studentRepo.GetAll()
}

func (u *studentUsecase) FindStudentById(id int) any {
	return u.studentRepo.GetById(id)
}

func (u *studentUsecase) Register(newStudent *model.Student) string {
	return u.studentRepo.Create(newStudent) + u.studentRepo.CreateUser(newStudent)
}

func (u *studentUsecase) Edit(student *model.Student) string {
	return u.studentRepo.Update(student)
}

func (u *studentUsecase) Unreg(id int) string {
	return u.studentRepo.Delete(id)
}

func (u *studentUsecase) UnregUser(username string) string {
	return u.studentRepo.DeleteUser(username)
}

func NewStudentUsecase(studentRepo repository.StudentRepo) StudentUsecase {
	return &studentUsecase{
		studentRepo: studentRepo,
	}
}
