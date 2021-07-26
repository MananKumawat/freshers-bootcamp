package Models

import (
	"awesomeProject/Day3/Exercise2/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
)
//GetAllUsers Fetch all user data
func GetAllStudents(student *[]Student) (err error) {
	if err = Config.DB.Preload(clause.Associations).Find(student).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateStudent(student *Student) (err error) {
	if err = Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}
//GetUserByID ... Fetch only one user by Id
func GetStudentByID(student *Student, id string) (err error) {
	if err = Config.DB.Preload(clause.Associations).Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}
//UpdateUser ... Update user
func UpdateStudent(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}
//DeleteUser ... Delete user
func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}