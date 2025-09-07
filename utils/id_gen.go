package utils

import (
	"github.com/bwmarrin/snowflake"
)

var (
	instance *snowflake.Node
)

func InitIdGeneratorClient() error {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	instance = node
	return nil
}

func GetId() int64 {
	return instance.Generate().Int64()
}
