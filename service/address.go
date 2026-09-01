package service

import (
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
)

// CreateAddress 创建地址信息函数
func CreateAddress(address *model.Address) error {
	err := dao.CreateAddress(address)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAddress 删除地址信息函数
func DeleteAddress(userId, addressId int64) error {
	err := dao.DeleteAddress(userId, addressId)
	if err != nil {
		return err
	}
	return nil
}

// UpdateAddress 修改地址信息
func UpdateAddress(userId int64, address *model.Address) error {
	err := dao.UpdateAddress(userId, address)
	if err != nil {
		return err
	}
	return nil
}
