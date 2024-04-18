package main

import (
	"fmt"
	"github.com/kongyixueyuan.com/kongyixueyuan/sdkInit"
	"github.com/kongyixueyuan.com/kongyixueyuan/service"
	"github.com/kongyixueyuan.com/kongyixueyuan/web"
	"github.com/kongyixueyuan.com/kongyixueyuan/web/controller"
	"os"
)

// 配置文件和初始化状态
const (
	configFile  = "config.yaml"
	initialized = false
	SimpleCC    = "simplecc"
)

func main() {
	// 构建信息结构体
	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/kongyixueyuan.com/kongyixueyuan/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     SimpleCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/kongyixueyuan.com/kongyixueyuan/chaincode/",
		UserName:        "User1",
	}
	//
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

	serviceSetup := service.ServiceSetup{
		ChaincodeID: SimpleCC,
		Client:      channelClient,
	}

	// 增加
	veh1 := service.Vehicle{
		VehicleID:  "1",
		OwnerName:  "yuenze",
		Signal:     "20",
		Reputation: "100",
		IsGood:     "true",
	}
	veh2 := service.Vehicle{
		VehicleID:  "2",
		OwnerName:  "gaolin",
		Signal:     "30",
		Reputation: "200",
		IsGood:     "true",
	}

	msg, err := serviceSetup.AddVehicle(veh1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add success", msg)
	}
	msg2, err := serviceSetup.AddVehicle(veh2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add success", msg2)
	}
	// query
	ans, err := serviceSetup.QueryByVehicleID("1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("query success", ans)
	}
	app := controller.Application{
		Fabric: &serviceSetup,
	}
	web.WebStart(&app)

}
