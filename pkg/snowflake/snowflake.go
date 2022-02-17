package snowflake

import (
	"go-scaffold/pkg/configs"
	"hash/crc64"
	"os"

	sf "github.com/bwmarrin/snowflake"
	"github.com/golang-module/carbon"
)

// 节点
var node *sf.Node

// Init 初始化雪花生成器
func Init() (err error) {
	// 时间戳：毫秒
	sf.Epoch = carbon.Parse(configs.AllConfig.Basic.OnlineTime).ToTimestampWithMillisecond()

	// 通过主机名获取唯一ID
	hostname, _ := os.Hostname()
	machineID := int64(crc64.Checksum([]byte(hostname), crc64.MakeTable(crc64.ISO)) % 1024)

	// 初始化节点
	node, err = sf.NewNode(machineID)
	if err != nil {
		return err
	}

	return nil
}

// GenerateID 生成分布式ID
func GenerateID() int64 {
	return node.Generate().Int64()
}
