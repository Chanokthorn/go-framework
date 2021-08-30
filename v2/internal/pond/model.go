package pond

import std_mysql "reflect-test/v2/internal/std/mysql"

type Pond struct {
	std_mysql.DBRootCommon
	PondID   *int    `db:"PondID" fake:"skip"`
	PondUUID *string `db:"PondUUID" fake:"skip"`
	Location *string `db:"Location" fake:"{city}"`
}

func (p *Pond) GetConfig() std_mysql.RootModelConfig {
	return std_mysql.RootModelConfig{
		TableName: "pond",
		IDField:   "PondID",
		UUIDField: "PondUUID",
	}
}
