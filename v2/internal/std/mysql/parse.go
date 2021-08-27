package std_mysql

import (
	"fmt"
	"strings"
)

func parseConfig(s string) (config RootModelConfig, err error) {
	//var config RootModelConfig
	splits := strings.Split(strings.ReplaceAll(s, " ", ""), ",")
	for _, split := range splits {
		innerSplits := strings.Split(split, ":")
		if len(innerSplits) != 2 {
			return RootModelConfig{}, fmt.Errorf(`invalid std config format: %s`, split)
		}
		k := innerSplits[0]
		v := innerSplits[1]
		switch k {
		case "tableName":
			config.TableName = v
		case "idField":
			config.IDField = v
		case "uuidField":
			config.UUIDField = v
		default:
			continue
		}

	}
	return config, nil
}
