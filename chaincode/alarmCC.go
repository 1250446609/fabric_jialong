package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"fmt"
	"bytes"
)

const DOC_TYPE = "alarmObj"

//保存alarm
// args:alarm
func PutAlarm(stub shim.ChaincodeStubInterface, alarm Alarm)([]byte, bool){

	alarm.ObjectType = DOC_TYPE
	b, err := json.Marshal(alarm)
	if err != nil {
		return nil, false
	}

	// 保存alarm 状态
	err = stub.PutState(alarm.Id, b)
	if err != nil {
		return nil, false
	}
	return b, true
}


// 添加信息
// args: alarmObject
// ID为 key, Alarm 为 value
func (t *AlarmChaincode) addAlarm(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2{
		return shim.Error("给定的参数个数不符合要求")
	}

	var alarm Alarm
	err := json.Unmarshal([]byte(args[0]), &alarm)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}
	// 查重: ID必须唯一 这里省略
	_, bl := PutAlarm(stub, alarm)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}