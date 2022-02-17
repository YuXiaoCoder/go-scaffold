package elasticsearch

import (
	"time"

	"github.com/olivere/elastic/v7"
)

func NewClient(endpoints []string) (client *elastic.Client, err error) {
	// 创建客户端
	client, err = elastic.NewClient(
		elastic.SetURL(endpoints...),                   // ES节点的URL端点
		elastic.SetSniff(false),                        // 关闭节点嗅探
		elastic.SetHealthcheckInterval(10*time.Second), // 调整节点健康检查间隔（60s => 10s）
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
