package main

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
