package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

// 新添加的字段是
// 链码id 链码gopath 链码路径  用户名
type InitInfo struct {
	ChannelID     string
	ChannelConfig string
	OrgAdmin      string
	OrgName       string
	OrdererOrgName	string
	OrgResMgmt *resmgmt.Client

	ChaincodeID string
	ChaincodeGoPath string
	ChaincodePath string
	UserName string
}