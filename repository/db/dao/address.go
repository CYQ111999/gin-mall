package dao

import (
	"gin-mall/pkg/e"
	"gin-mall/repository/db"
	"gin-mall/repository/db/model"
	"time"
)

// CreateAddress 地址创建函数
func CreateAddress(address *model.Address) error {
	err := db.DB.Create(address).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteAddress 地址删除函数
func DeleteAddress(userId, addressId int64) error {
	// 根据用户id和地址id删除信息,
	result := db.DB.Where("user_id = ? and address_id = ?", userId, addressId).
		Delete(&model.Address{})
	// 如果是系统错误就返回错误信息
	if result.Error != nil {
		return result.Error
	}
	// 如果修改了0行数据说明地址或者用户id有错，返回地址不存在的自定义错误信息
	if result.RowsAffected == 0 {
		return e.NewError(e.ERROR_ADDRESS_NOT_EXIST)
	}
	return nil
}

// UpdateAddress 修改地址信息
func UpdateAddress(userId int64, address *model.Address) error {
	// 修改地址信息
	result := db.DB.Model(&model.Address{}).
		Where("user_id = ? and address_id = ?", userId, address.AddressID).
		Updates(map[string]interface{}{
			"name":       address.Name,
			"phone":      address.Phone,
			"address":    address.Address,
			"updated_at": time.Now(),
		})
	// 如果是系统错误就返回错误信息
	if result.Error != nil {
		return result.Error
	}
	// 如果修改了0行数据说明地址或者用户id有错，返回地址不存在的自定义错误信息
	if result.RowsAffected == 0 {
		return e.NewError(e.ERROR_ADDRESS_NOT_EXIST)
	}
	return nil
}

// GetAddressById 通过地址id获取地址信息
func GetAddressById(id int64) (*model.Address, error) {
	address := new(model.Address)
	// 查到第一个相同地址id的地址就返回地址信息
	err := db.DB.Model(&model.Address{}).
		Where("address_id = ?", id).
		First(address).Error
	if err != nil {
		return nil, err
	}
	return address, nil
}
