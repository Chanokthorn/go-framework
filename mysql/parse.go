package mysql

import (
	"fmt"
	"strings"
)

func parseConfig(s string) (config StdConfig, err error) {
	//var config StdConfig
	splits := strings.Split(strings.ReplaceAll(s, " ", ""), ",")
	for _, split := range splits {
		innerSplits := strings.Split(split, ":")
		if len(innerSplits) != 2 {
			return StdConfig{}, fmt.Errorf(`invalid std config format: %s`, split)
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
		case "recursiveOnGetAll":
			if v == "false" {
				config.RecursiveOnGetAll = false
			}
			if v == "true" {
				config.RecursiveOnGetAll = true
			}
		default:
			continue
		}

	}
	return config, nil
}
