package main

import(
	"log"
	"../../model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
			Username:"linsan",
			PasswordHash:model.GeneratePasswordHash("abc123"),
			Posts:[]model.Post{
				{Body:"Beautiful day in Portland"},
			},
		},
		{
			Username:"rene",
			PasswordHash:model.GeneratePasswordHash("abc123"),
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
