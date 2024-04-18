package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// 保存一个车的记录
func PutVehicle(stub shim.ChaincodeStubInterface, veh Vehicle) ([]byte, bool) {

	b, err := json.Marshal(veh)
	if err != nil {
		return nil, false
	}

	// 保存veh状态
	err = stub.PutState(veh.VehicleID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func GetVehicleInfo(stub shim.ChaincodeStubInterface, vehicleID string) (Vehicle, bool) {
	var veh Vehicle
	// 根据vehicleID查询信息状态
	b, err := stub.GetState(vehicleID)
	if err != nil {
		return veh, false
	}

	if b == nil {
		return veh, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &veh)
	if err != nil {
		return veh, false
	}

	// 返回结果
	return veh, true
}

// 增加一个车辆的信息，第一个参数是车辆信息，直接反序列化到veh，第二个参数是
func (t *SimpleChaincode) addVehicle(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var veh Vehicle
	err := json.Unmarshal([]byte(args[0]), &veh)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: 车辆ID必须唯一
	_, exist := GetVehicleInfo(stub, veh.VehicleID)
	if exist {
		return shim.Error("要添加的车辆ID已存在")
	}

	_, bl := PutVehicle(stub, veh)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根据vehicleID进查询
func (t *SimpleChaincode) queryVehicleByVehicleID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据vehicleID查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据vehicleID查询信息失败")
	}

	if b == nil {
		return shim.Error("根据vehicleID没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var veh Vehicle
	err = json.Unmarshal(b, &veh)
	if err != nil {
		return shim.Error("反序列化veh信息失败")
	}

	// 返回
	result, err := json.Marshal(veh)
	if err != nil {
		return shim.Error("序列化veh信息时发生错误")
	}
	return shim.Success(result)
}

// 根据vehicleID更新信息
// args: 第一个参数是一个完整的vehicle，第二个参数作用未知
func (t *SimpleChaincode) updateVehicle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 反序列化第一个参数到veh里
	var veh Vehicle
	err := json.Unmarshal([]byte(args[0]), &veh)
	if err != nil {
		return shim.Error("反序列化veh信息失败")
	}

	// 根据vehicleID查询信息
	result, bl := GetVehicleInfo(stub, veh.VehicleID)
	if !bl {
		return shim.Error("根据vehicleID查询信息时发生错误")
	}

	result.VehicleID = veh.VehicleID
	result.OwnerName = veh.OwnerName
	result.Signal = veh.Signal
	result.Reputation = veh.Reputation
	result.IsGood = veh.IsGood

	_, bl = PutVehicle(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

// 根据vehicleID删除信息
// args: entityID，第二个未知
func (t *SimpleChaincode) delVehicle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	/*var edu Education
	result, bl := GetEduInfo(stub, info.EntityID)
	err := json.Unmarshal(result, &edu)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("车辆信息删除成功"))
}
