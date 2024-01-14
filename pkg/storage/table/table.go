package table

import "github.com/forg002-ctrl/managment_system/pkg/storage"

type Table struct {
	Name string
	Attributes map[string]string
}

func (t *Table) InitTable(s *storage.PostgresClient) error {
	return s.CreateTable(t.Name, t.Attributes);
}