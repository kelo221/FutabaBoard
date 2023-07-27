package Database

import (
	DataModels "backend/ORM"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Init() {

	var err error

	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&DataModels.Post{}, &DataModels.Thread{}, &DataModels.Bans{})
	if err != nil {
		return
	}

	/*	db.Create(&DataModels.Thread{
			Name:  "Anon 2",
			Text:  "Hello I'm OP 2",
			Flags: "US",
		})

		db.Create(&DataModels.Thread{
			Name:  "Anon",
			Text:  "Hello I'm OP",
			Flags: "US",
		})

		db.Create(&DataModels.Post{
			UnixTime:     time.Now(),
			Name:         "Anon",
			Text:         "Hello World1",
			Flags:        "FI",
			ParentThread: 1,
			PostImage:    DataModels.PostImage{},
		})

		db.Create(&DataModels.Post{
			UnixTime:     time.Now(),
			Name:         "Anon",
			Text:         "Hello World2",
			Flags:        "US",
			ParentThread: 1,
			PostImage:    DataModels.PostImage{},
		})*/

}

/*	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&DataModels.Post{}, &DataModels.Thread{})
	if err != nil {
		return
	}

	db.Create(&DataModels.Thread{
		Text: "thread1",
	})
	db.Create(&DataModels.Post{
		Text:     "Post1",
		ThreadID: 1,
	})

	db.Create(&DataModels.Thread{
		Text: "thread2",
	})

	db.Create(&DataModels.Post{
		Text:     "Post2",
		ThreadID: 2,
	})*/

// Create

// Read
//var user DataModels.Account
////db.First(&user, 0) // find user with integer primary key
//db.First(&user, "Username = ?", "Test User") // find user with code D42
//fmt.Println("User ID:", user.Password)
//fmt.Println("Username:", user.Username)

//
//// Update - update user's price to 200
//db.Model(&user).Update("Username", "changed value")
//// Update - update multiple fields
//db.Model(&user).Updates(DataModels.Account{
//	Username: "Updated Again",
//	Password: "Updated Again",
//}) // non-zero fields
//db.Model(&user).Updates(map[string]interface{}{"Username": "Hello", "Password": "F42"})
//
//// Delete - delete user
//db.Delete(&user, 1)
