package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct{}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获得参数
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addVehicle" {
		return t.addVehicle(stub, args)
	} else if fun == "queryVehicleByVehicleID" {
		return t.queryVehicleByVehicleID(stub, args)
	} else if fun == "updateVehicle" {
		return t.updateVehicle(stub, args)
	} else if fun == "delVehicle" {
		return t.delVehicle(stub, args)
	}

	return shim.Error("指定的函数名称错误")
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("启动SimpleChaincode时发生错误: %s", err)
	}
}
