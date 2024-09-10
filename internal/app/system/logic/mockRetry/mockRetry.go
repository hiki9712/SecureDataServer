package mockRetry

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func init() {
	service.RegisterMockRetry(New())
}

type SMockRetry struct {
}

func New() *SMockRetry {
	return &SMockRetry{}
}

func (s *SMockRetry) ResolveReq(ctx context.Context, req *system.MockRetryReq) (data g.Map, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		reqJson, _ := json.Marshal(req)
		err = json.Unmarshal(reqJson, &data)
	})
	return
}

func (s *SMockRetry) FetchTableByTaskID(ctx context.Context, data g.Map) (tableData *model.Table, err error) {
	var (
		TaskID    int64
		DB        string
		Table     string
		TaskInfos []*entity.TaskInfo
		TaskData  g.Map
		Infos     []*entity.ElectricityInfo
	)
	err = g.Try(ctx, func(ctx context.Context) {
		tableData = &model.Table{}
		TaskID = gconv.Int64(data["taskID"])
		err = g.DB("default").Model("task").Ctx(ctx).Where("task_id=?", TaskID).Scan(&TaskInfos)
		g.Log().Info(ctx, "taskInfo:", TaskInfos[0])
		taskJson, _ := json.Marshal(TaskInfos[0])
		err = json.Unmarshal(taskJson, &TaskData)
		DB = gconv.String(TaskData["db_name"])
		Table = gconv.String(TaskData["table_name"])
		fmt.Println(TaskID, DB, Table)
		err = g.DB(DB).Model(Table).Ctx(ctx).Scan(&Infos)
		tableData.TableContent = Infos
		tableData.TableName = Table
		tableData.TaskID = TaskID
	})
	return
}

func (s *SMockRetry) SendToKafka(ctx context.Context, tableData *model.Table) (err error) {
	//kafka配置
	kafkaConfig := g.Cfg().MustGet(ctx, "kafka").Map()
	brokers := kafkaConfig["brokers"].([]interface{})
	brokerList := make([]string, len(brokers))
	for i, broker := range brokers {
		brokerList[i] = broker.(string)
	}
	topic := kafkaConfig["topic"].(string)
	producerConfig := kafkaConfig["producer"].(map[string]interface{})
	return_successes := producerConfig["return_successes"].(bool)
	//required_acks := producerConfig["required_acks"].(int)
	//retry_max := producerConfig["retry_max"].(int)

	config := sarama.NewConfig()
	config.Producer.Return.Successes = return_successes
	config.Producer.RequiredAcks = sarama.RequiredAcks(-1)
	config.Producer.Retry.Max = 1

	//初始化Kafka生产者
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		g.Log().Error(ctx, "sarama.NewSyncProducer err:", err)
		return
	}
	defer producer.Close()

	messageData, err := json.Marshal(tableData)
	if err != nil {
		g.Log().Error(ctx, "json.Marshal err:", err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(messageData),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		g.Log().Error(ctx, "SendMessage err:", err)
		return
	}
	g.Log().Info(ctx, "partition:", partition, "offset:", offset)
	return
}
