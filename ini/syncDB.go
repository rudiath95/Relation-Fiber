package ini

import (
	"fmt"

	"github.com/rudiath95/RelationFiber/models"
)

func SyncDatabases() {
	err := DB.AutoMigrate(
		models.User{},
		models.Locker{},
		models.Post{},
		models.Tag{},
	)

	if err != nil {
		fmt.Println("Failed to migrate database")
	} else {
		fmt.Println("Successfully migrated database")
	}
}
