package repositories

import (
	"gorm.io/gorm"
)

type AuthorizationRepository struct {
	DB *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) *AuthorizationRepository {
	return &AuthorizationRepository{DB: db}
}

func (r *AuthorizationRepository) UserHasPermission(email string, permissionID int) (bool, error) {
	var count int64
	err := r.DB.Table("users").
		Joins("JOIN user_types ON users.user_type_id = user_types.id").
		Joins("JOIN user_type_has_role ON user_types.id = user_type_has_role.user_type_id").
		Joins("JOIN roles ON user_type_has_role.role_id = roles.id").
		Joins("JOIN role_permission ON roles.id = role_permission.role_id").
		Joins("JOIN permissions ON role_permission.permission_id = permissions.id").
		Where("users.email = ? AND permissions.id = ?", email, permissionID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
