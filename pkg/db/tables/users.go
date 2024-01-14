package tables

import storage "github.com/forg002-ctrl/managment_system/pkg/storage/table"

func GetUserTable() *storage.Table{
	return &storage.Table{
		Name: "users",
		Attributes: map[string]string{
			"user_id":  "SERIAL PRIMARY KEY",
			"username": "VARCHAR(50) NOT NULL",
			"email":    "VARCHAR(100) NOT NULL",
			"password": "VARCHAR(255) NOT NULL",
		},
	}
}