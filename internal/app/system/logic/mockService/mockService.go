package mockService

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func init() {
	service.RegisterMockService(New())
}

type SMockService struct {
}

func New() *SMockService {
	return &SMockService{}
}

func (s *SMockService) ResolveReq(ctx context.Context, req *system.MockServiceReq) (data g.Map, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		reqJson, _ := json.Marshal(req)
		err = json.Unmarshal(reqJson, &data)
	})
	return
}

func (s *SMockService) FetchTable(ctx context.Context, data g.Map) (tableData *model.Table, err error) {
	var (
		TaskID int64
		DB     string
		Table  string
		//Infos  []*entity.ElectricityInfo
	)
	err = g.Try(ctx, func(ctx context.Context) {
		tableData = &model.Table{}
		TaskID = gconv.Int64(data["taskID"])
		DB = gconv.String(data["db"])
		Table = gconv.String(data["table"])
		fmt.Println(TaskID, DB, Table)
		//err = g.DB(DB).Model(Table).Ctx(ctx).Scan(&Infos)
		//tableData.TableContent = Infos
		tableData.TableName = Table
		tableData.TaskID = TaskID
	})
	return
}

func (s *SMockService) SendToKafka(ctx context.Context, tableData *model.Table) (err error) {
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
