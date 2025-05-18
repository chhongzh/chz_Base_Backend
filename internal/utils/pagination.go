package utils

import "gorm.io/gorm"

// 这个函数将会包装一些分页的逻辑
func Pagination(tx *gorm.DB, currentPage int, itemsPerPage int) *gorm.DB {
	return tx.Offset((currentPage - 1) * itemsPerPage).Limit(itemsPerPage)
}
