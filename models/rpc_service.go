package models

import (
	"fmt"
	"gateway-camille/db"
)

type RpcService struct {
	ID          int    `json:"id" gorm:"column:id;primary_key"`
	Route       string `json:"route" gorm:"column:route"`
	Port        string `json:"port" gorm:"column:port"`
	RpcHost     string `json:"rpcHost" gorm:"column:rpcHost"`
	Weight      string `json:"weight" gorm:"column:weight"`
	BanIP       string `json:"banIP" gorm:"column:banIP"`
	BalanceType int    `json:"balanceType" gorm:"column:balanceType"`
	WhiteIP     string `json:"whiteIP" gorm:"column:whiteIP"`
	RsaFlag     int    `json:"rsaFlag" gorm:"column:rsaFlag"`
}

// CreateRpcService 创建一个新的 RpcService 记录
func CreateRpcService(rpcService RpcService) error {
	dbConnection := db.GetGiftDBConnection()
	result := dbConnection.Create(&rpcService)
	if result.Error != nil {
		return fmt.Errorf("Failed to create RpcService! err: %v", result.Error)
	} else {
		fmt.Printf("Created RpcService: %+v\n", rpcService)
	}
	return nil
}

// ReadNodesRpcService 根据 rpc 的 port 获取节点
func ReadNodesRpcService(port string) (RpcService, error) {
	var rpcServices RpcService
	dbConnection := db.GetGiftDBConnection()
	result := dbConnection.Where("port = ?", port).First(&rpcServices)
	if result.Error != nil {
		return RpcService{}, fmt.Errorf("Failed to read RpcServices! err: %v", result.Error)
	} else if result.RowsAffected == 0 {
		return RpcService{}, fmt.Errorf("No RpcServices found for rpcHost: %s", port)
	} else {
		fmt.Printf("Read RpcServices: %+v\n", rpcServices)
	}
	return rpcServices, nil
}

// ReadHealthyNodesRpcService 获取健康节点
func ReadHealthyNodesRpcService(port string) ([]RpcService, error) {
	var rpcServices []RpcService
	dbConnection := db.GetGiftDBConnection()
	result := dbConnection.Where("port = ? AND rsa_flag = 0", port).Find(&rpcServices)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to read Healthy RpcServices! err: %v", result.Error)
	} else if result.RowsAffected == 0 {
		return nil, fmt.Errorf("No Healthy RpcServices found for port: %s", port)
	} else {
		fmt.Printf("Read Healthy RpcServices: %+v\n", rpcServices)
	}
	return rpcServices, nil
}

// UpdateRpcService 更新一个已存在的 RpcService 记录
func UpdateRpcService(rpcHost string, rsaFlag int) error {
	dbConnection := db.GetGiftDBConnection()
	rpcService := RpcService{RpcHost: rpcHost}
	result := dbConnection.Model(&rpcService).Update("rsa_flag", rsaFlag)
	if result.Error != nil {
		return fmt.Errorf("Failed to update RpcService! err: %v", result.Error)
	} else {
		fmt.Printf("Updated RpcService: %+v\n", rpcService)
	}
	return nil
}

// DeleteRpcService 删除一个 RpcService 记录
func DeleteRpcService(port string) error {
	dbConnection := db.GetGiftDBConnection()
	result := dbConnection.Where("port = ?", port).Delete(RpcService{})
	if result.Error != nil {
		return fmt.Errorf("Failed to delete RpcService! err: %v", result.Error)
	} else if result.RowsAffected == 0 {
		return fmt.Errorf("No RpcService found for port: %s", port)
	} else {
		fmt.Printf("Deleted RpcService port: %s\n", port)
	}
	return nil
}
