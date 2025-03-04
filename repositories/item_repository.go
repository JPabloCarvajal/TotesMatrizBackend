package repositories

import (
	"totesbackend/models"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) GetItemByID(id string) (*models.Item, error) {
	var item models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) SearchItemsByID(query string) ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").
		Where("id LIKE ?", query+"%").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) SearchItemsByName(query string) ([]models.Item, error) {
	var items []models.Item
	err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").
		Where("LOWER(name) LIKE LOWER(?)", query+"%").
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemRepository) UpdateItemState(id string, state bool) (*models.Item, error) {
	var item models.Item
	if err := r.DB.Preload("ItemType").Preload("AdditionalExpenses").First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}

	item.ItemState = state

	if err := r.DB.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) UpdateItem(item *models.Item) error {
	// Buscar el item en la base de datos
	var existingItem models.Item
	if err := r.DB.First(&existingItem, "id = ?", item.ID).Error; err != nil {
		return err // Retorna error si no se encuentra
	}

	// Actualizar el item con los nuevos valores
	if err := r.DB.Save(item).Error; err != nil {
		return err // Retorna error si falla la actualización
	}

	return nil // Éxito
}
