package db

import (
	"fmt"

	"github.com/forg002-ctrl/managment_system/pkg/db/tables"
	"github.com/forg002-ctrl/managment_system/pkg/storage"
)

func InitTables(s *storage.PostgresClient) error {
	if s == nil {
		return fmt.Errorf("Error: PostgresClient is nil")
	}
	err := tables.GetUserTable().InitTable(s);
	if err != nil {
		return fmt.Errorf("Error: InitTable failed - ", err)
	}

	return nil
}