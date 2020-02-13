package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) SaveAlarm(alarm Alarm) (string, error) {
	eventID := "eventAddAlarm"
	reg, notifier = regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	//将alarm对象序列化为字节数组
	b, err = json.Marshal(alarm)
	if err != nil {
		return "", fmt.Errorf("指定的alarm对象序列化时候发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addAlarm", Args: [][]byte{b, []byte(eventID)}}
	response, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(response.TransactionID), nil
}
