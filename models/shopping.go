package models

import (
	"gorm.io/gorm"
)

type Shopping struct{
	
	gorm.Model
	Id       int     `form:"id" json:"id" validate:"required"`
	Name     string  `form:"name" json:"name" validate:"required"`
	Image	 string  `form:"image" json:"image"`
	Quantity int     `form:"quantity" json:"quantity" validate:"required"`
	Price    float32 `form:"price" json:"price" validate:"required"`
}



// CRUD
func CreateShopping(db *gorm.DB, newShopping *Shopping) (err error) {
	err = db.Create(newShopping).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadShoppings(db *gorm.DB, shoppings *[]Shopping)(err error) {
	err = db.Find(shoppings).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadShoppingById(db *gorm.DB, shopping *Shopping, id int)(err error) {
	err = db.Where("id=?", id).First(shopping).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateShopping(db *gorm.DB, shopping *Shopping)(err error) {
	db.Save(shopping)

	return nil
}
func DeleteShoppingById(db *gorm.DB, shopping *Shopping, id int)(err error) {
	db.Where("id=?", id).Delete(shopping)
 
	return nil
}


