package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)

// 车辆ID 车主姓名 信号 信誉值 是否良性

type Vehicle struct {
	VehicleID  string `json:"vehicleID"`
	OwnerName  string `json:"ownerName"`
	Signal     string `json:"signal"`
	Reputation string `json:"reputation"`
	IsGood     string `json:"isGood"`
	Records    []Record
}

// 成本 奖励
type Record struct {
	cost   string `json:"cost"`
	reward string `json:"reward"`
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

// 必须注册链码事件
func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

//ServiceSetup 事件结果
func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
