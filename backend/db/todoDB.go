package db

import (
	"gorm.io/gorm"
	"fmt"
	// "encoding/json"
)

type UserDB struct {
	Email string `gorm:"primaryKey" json:"email"`
	Name string   `json:"name"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"createdat"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updatedat"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedat"`
}

type TodoDB struct {
	ID uint64 `gorm:"primaryKey`
	UserEmail string  `json:"useremail"`
	UserDBs []UserDB `gorm:"foreignKey:Email;references:UserEmail"`
	Description string  `json:"description"`
	Complete bool `json:"complete"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"createdat"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updatedat"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedat"`
}

func ReadAll() (*[]UserDB, error){
	db := DB()
	// var userDB UserDB
	// var usersDB UsersDB
	users := []UserDB{}
	db.Find(&users)
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	fmt.Print(&users)

	// res2B, _ := json.Marshal(users)
	// fmt.Println(string(res2B))
	// usersDB = users
	// // db.Find(&users)
	// fmt.Print(&usersDB)
	return &users,nil
}

func Read(code string) (*UserDB, error) {
	db := DB()
	
	var userDB UserDB
	if err := db.First(&userDB, "email = ?",code).Error; err != nil {
		return nil, err
	}
	fmt.Print(&userDB)
	return &userDB, nil
}

func Create(user UserDB) error {
	db := DB()
	userData := UserDB{
		// Code: user.Code,
		// Name: user.Name,
		Email: user.Email,
	}
	if err := db.Create(&userData).Error; err != nil {
		return err
	}
	return nil
}

func ReadTodo(userEmail string) (*[]TodoDB, error) {
	db := DB()
	
	todoDB := []TodoDB{}
	// if err := db.First(&todoDB,"user_email = ?",userEmail).Error; err != nil {
	// 	return nil, err
	// }
	if err := db.Where("user_email = ?", userEmail).Find(&todoDB).Error; err != nil {
		return nil, err
	}
	fmt.Print(&todoDB)
	return &todoDB, nil
}

func UpdateTodo(todo TodoDB)  error {
	db := DB()
	
	todoDB := TodoDB{
		ID:todo.ID,
		Description : todo.Description,
		UserEmail : todo.UserEmail,
		Complete : todo.Complete,
	}
	
	if err := db.Model(&todoDB).Update("complete", todo.Complete).Error; err != nil {
		return err
	}
	fmt.Print(&todoDB)
	return nil
}

func DeleteTodo(todo TodoDB)  error {
	db := DB()
	
	todoDB := TodoDB{
		ID:todo.ID,
		Description : todo.Description,
		UserEmail : todo.UserEmail,
		Complete : todo.Complete,
	}
	
	if err := db.Delete(&todoDB).Error; err != nil {
		return err
	}
	fmt.Print(&todoDB)
	return nil
}


func CreateTodo(todo TodoDB) error {
	db := DB()
	userData := TodoDB{
		Complete : todo.Complete,
		Description : todo.Description,
		UserEmail: todo.UserEmail,
	}

	if err := db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
}