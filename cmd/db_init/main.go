package main

import(
	"log"
	"../../model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

func main(){
	log.Println("DB Init...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	
	db.DropTableIfExists(model.User{},model.Post{})
	db.CreateTable(model.User{},model.Post{})

	user := []model.User{
		{
			Username:     "linsan",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "i@bonfy.im",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("i@bonfy.im")),
			Posts:[]model.Post{
				{Body:"Beautiful day in Portland"},
			},
		},
		{
			Username:     "rene",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "rene@test.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("rene@test.com")),
			Posts:[]model.Post{
				{Body:"The Avengers movie was so cool!"},
				{Body:"Sun shine is beautiful"},
			},
		},
	}

	for _,u := range user{
		db.Debug().Create(&u)
	}
}
