package migrations

import (
	"fmt"
	"restapi-users-management/database"
	"restapi-users-management/models"
)

func MigrateTables() {
	migrateUserTable()
	migrateSessionTable()
}

func migrateUserTable() {

	// database.DB.Migrator().DropTable(models.User{})

	if !database.DB.Migrator().HasTable(models.User{}) {
		err := database.DB.AutoMigrate(models.User{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "User")
	}
}

func migrateSessionTable() {

	// database.DB.Migrator().DropTable(models.Session{})

	if !database.DB.Migrator().HasTable(models.Session{}) {
		err := database.DB.AutoMigrate(models.Session{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Session")
	}
}
