package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

//雪花算法

var node *snowflake.Node

func Init(startTime string, format string, machineID int64) error {
	st, err := time.Parse(format, startTime)
	if err != nil {
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return nil
}

func GetID() int64 {
	return node.Generate().Int64()
}
