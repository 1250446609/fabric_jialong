package main

import (
	"os"
	"fmt"
	"github.com/kongyixueyuan.com/alarm/sdkInit"

	"github.com/kongyixueyuan.com/alarm/service"

	"encoding/json"
	"github.com/kongyixueyuan.com/alarm/web/controller"
	"github.com/kongyixueyuan.com/alarm/web"
)

const (
	configFile = "config.yaml"
	initialized = false
	SimpleCC = "AlarmCC"
)

func main()  {
	initInfo := &sdkInit.InitInfo{

		ChannelID: "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/kongyixueyuan.com/alarm/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID: SimpleCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/kongyixueyuan.com/alarm/chaincode/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)


	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID: "AlarmCC",
		Client:channelClient,
	}

	alarm := service.Alarm{
		Id: "张三",
		SiteId: "徐州站房",
		NoticeType: "报警通知",
		AlarmDetail: "这里是报警详情",
		AlarmLevel: "三级报警",
		SiteType: "配电房",
		TriggerValue: "100度阈值",
		AlarmParam: "湿度",
		AlarmTime: "报警时间",
		Auditor: "作者",
	}


	msg, err := serviceSetup.SaveAlarm(alarm)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}

	//===========================================//

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)
}