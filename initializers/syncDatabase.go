package initializers

import "github.com/PayamAnari/kubespaces/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
