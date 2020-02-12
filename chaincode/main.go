package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
)

type AlarmChaincode struct {

}

func (t *AlarmChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response  {
	return shim.Success(nil)
}

func (t *AlarmChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response  {
	//获取用户意图
	fun, args = stub.GetFunctionAndParameters()

	if fun == "addAlarm"{
		return t.addAlarm(stub, args)
	}

}

func main(){
	err := shim.Start(new(AlarmChaincode))
	if err != nil {
		fmt.Printf("启动AlarmChaincode时候发生错误:%s", err)
	}
}