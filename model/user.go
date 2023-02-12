package model

import (
	"fmt"
	"meetout-ecr/libs"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Tablename string
	User      User
}

type Role struct {
	Id   string `bson:"id"`
	Name string `bson:"name"`
}

type User struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name"  json:"last_name" `
	Username  string `bson:"username"   json:"username"`
	Password  string `bson:"password"   json:"password"`
	Role      Role   `bson:"role"       json:"role"`
}

/********For User Model Methods ************/

func (usermodel UserModel) Insert(dao libs.DAO) (*mongo.InsertOneResult, error) {
	fmt.Println(usermodel)
	return dao.GetDB().Collection(usermodel.Tablename).InsertOne(dao.Ctx, usermodel.User)
}

func (usermodel UserModel) GetAllUsers(dao libs.DAO) (*[]User, error) {
	cur, err := dao.GetDB().Collection(usermodel.Tablename).Find(dao.Ctx, struct{}{})
	if err != nil {
		return nil, err
	}
	var results *[]User = &[]User{}
	err = cur.All(dao.Ctx, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

/********For User Object Methods ************/

func (user User) GetName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) SetName(first_name, last_name string) {
	user.FirstName = first_name
	user.LastName = last_name
}
