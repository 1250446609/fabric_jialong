package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

type Alarm struct {
	ObjectType  string `json:"docType"`
	Id          string `json:"Id"`          //站房
	SiteId      string `json:"SiteId"`      //站房
	NoticeType  string `json:"NoticeType"`  //通知类型
	AlarmDetail string `json:"AlarmDetail"` //报警详情

	AlarmLevel string `json:"AlarmLevel"` //报警级别

	SiteType string `json:"SiteType"` //站房类型

	TriggerValue string        `json:"TriggerValue"` //触发值
	AlarmParam   string        `json:"AlarmParam"`   //报警参数
	AlarmTime    string        `json:"AlarmTime"`    //报警时间
	Auditor      string        `json:"Auditor"`      //作者
	Historys     []HistoryItem `json:"alarmParam"`
}

type HistoryItem struct {
	TxId  string
	Alarm Alarm
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
