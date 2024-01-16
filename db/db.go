package database

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbName     = "postgres"
	dbUser     = "postgres"
	dbPassword = "123"
)

//	func ConnectDB() (db *gorm.DB, err error) {
//		dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
//			dbHost, dbPort, dbName, dbUser, dbPassword)
//		db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
//		if err != nil {
//			return
//		}
//		db.AutoMigrate(&models.User{})
//		return
//	}
// func ConnectDB() (db *gorm.DB, err error) {
// 	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
// 		dbHost, dbPort, dbName, dbUser, dbPassword)
// 	db, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
// 	if err != nil {
// 		return
// 	}

// 	err = db.AutoMigrate(&Users{})
// 	if err != nil {
// 		return
// 	}
// 	return
// }
