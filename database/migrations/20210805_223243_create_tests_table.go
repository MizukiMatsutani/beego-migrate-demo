package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateTestsTable_20210805_223243 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTestsTable_20210805_223243{}
	m.Created = "20210805_223243"

	migration.Register("CreateTestsTable_20210805_223243", m)
}

// Run the migrations
func (m *CreateTestsTable_20210805_223243) Up() {
	m.SQL(`
	create table tests (
		id bigint auto_increment primary key
		, name varchar(20) not null
	)
	`)
}

// Reverse the migrations
func (m *CreateTestsTable_20210805_223243) Down() {
	m.SQL("DROP TABLE if exists tests")
}
