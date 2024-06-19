package gorm

import "api-project/internal/domain/models"

func (r *repository) DeleteUser(id *uint) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
