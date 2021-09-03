package pond

import std_mysql "reflect-test/v3/internal/lib/mysql"

type Pond struct {
	std_mysql.RootCommon
	PondID   *int    `db:"PondID" fake:"skip"`
	PondUUID *string `db:"PondUUID" fake:"skip"`
	Location *string `db:"Location" fake:"{country}"`
}

func (p *Pond) GetConfig() std_mysql.RootConfig {
	return std_mysql.RootConfig{
		TableName: "pond",
		IDField:   "PondID",
		UUIDField: "PondUUID",
	}
}
