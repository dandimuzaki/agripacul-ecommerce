package repository

import (
	"context"
	"debian-ecommerce/internal/data/entity"
	infra "debian-ecommerce/internal/infra/transaction"
	"debian-ecommerce/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartByCustomerID(ctx context.Context, customerID uint) (*entity.Cart, error)
	GetSelectedCartItems(ctx context.Context, customerID uint) ([]entity.CartItem, error)
	CreateCart(ctx context.Context, cart *entity.Cart) error
	AddItem(ctx context.Context, item *entity.CartItem) error
	GetCartItemBySKU(ctx context.Context, cartID, skuID uint) (*entity.CartItem, error)
	BatchRemoveItems(ctx context.Context, itemIDs []uint) error
	ClearCart(ctx context.Context, cartID uint) error
	DeleteCart(ctx context.Context, cartID uint) error
	UpdateItem(ctx context.Context, itemID, customerID uint, data map[string]interface{}) error
	GetItemByID(ctx context.Context, itemID, customerID uint) (*entity.CartItem, error)
	BatchSelectItem(ctx context.Context, customerID uint, data map[string]interface{}) error
}

type cartRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCartRepository(db *gorm.DB, log *zap.Logger) CartRepository {
	return &cartRepository{
		db:  db,
		log: log,
	}
}

func (r *cartRepository) GetCartByCustomerID(ctx context.Context, customerID uint) (*entity.Cart, error) {
	db := infra.GetDB(ctx, r.db)
	var cart entity.Cart
	err := db.Preload("Items", "quantity > ?", 0).
		Preload("Items.SKU", "status = ?", entity.SKUStatusActive).
		Preload("Items.SKU.Product").
		Preload("Items.SKU.Images").
		Preload("Items.SKU.SKUVariantValues.VariantValue").
		Preload("Items.SKU.SKUVariantValues.VariantValue.VariantType").
		Where("customer_id = ?", customerID).
		First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *cartRepository) GetSelectedCartItems(ctx context.Context, customerID uint) ([]entity.CartItem, error) {
	db := infra.GetDB(ctx, r.db)
	var cartItems []entity.CartItem
	err := db.Model(&entity.CartItem{}).
	Preload("SKU", "status = ?", entity.SKUStatusActive).
	Preload("SKU.Product").
	Preload("SKU.Images").
	Preload("SKU.SKUVariantValues.VariantValue").
	Preload("SKU.SKUVariantValues.VariantValue.VariantType").
	Where("is_selected = ?", true).Where("quantity > ?", 0).Find(&cartItems).Error
	if err != nil {
		r.log.Error("Failed to get selected cart items", zap.Error(err))
		return nil, err
	}
	return cartItems, nil
}

func (r *cartRepository) CreateCart(ctx context.Context, cart *entity.Cart) error {
	db := infra.GetDB(ctx, r.db)
	return db.Create(cart).Error
}

func (r *cartRepository) AddItem(ctx context.Context, item *entity.CartItem) error {
	db := infra.GetDB(ctx, r.db)
	return db.Create(item).Error
}

func (r *cartRepository) GetCartItemBySKU(ctx context.Context, cartID, skuID uint) (*entity.CartItem, error) {
	db := infra.GetDB(ctx, r.db)
	var item entity.CartItem
	err := db.Where("cart_id = ? AND sku_id = ?", cartID, skuID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *cartRepository) BatchRemoveItems(ctx context.Context, itemIDs []uint) error {
	db := infra.GetDB(ctx, r.db)
	return db.Model(&entity.CartItem{}).Where("cart_items.id IN ?", itemIDs).Updates(map[string]interface{}{
		"quantity": 0,
		"is_selected": false,
	}).Error
}

func (r *cartRepository) ClearCart(ctx context.Context, cartID uint) error {
	db := infra.GetDB(ctx, r.db)
	return db.Model(&entity.CartItem{}).Where("cart_id = ?", cartID).Updates(map[string]interface{}{
		"quantity": 0,
		"is_selected": false,
	}).Error
}

func (r *cartRepository) DeleteCart(ctx context.Context, cartID uint) error {
	db := infra.GetDB(ctx, r.db)
	return db.Delete(&entity.Cart{}, cartID).Error
}

func (r *cartRepository) UpdateItem(ctx context.Context, itemID, customerID uint, data map[string]interface{}) error {
	db := infra.GetDB(ctx, r.db)

	subQuery := db.Model(&entity.Cart{}).
		Select("id").
		Where("customer_id = ?", customerID)

	result := db.Model(&entity.CartItem{}).
		Where("id = ?", itemID).
		Where("cart_id IN (?)", subQuery). // Use IN for subqueries to be safe
		Updates(data)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrCartItemNotFound
	}

	return nil
}

func (r *cartRepository) GetItemByID(ctx context.Context, itemID, customerID uint) (*entity.CartItem, error) {
	db := infra.GetDB(ctx, r.db)
	var item entity.CartItem

	subQuery := db.Model(&entity.Cart{}).
		Select("id").
		Where("customer_id = ?", customerID)

	err := db.Model(&entity.CartItem{}).
		Where("id = ?", itemID).
		Where("cart_id IN (?)", subQuery).First(&item).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *cartRepository) BatchSelectItem(ctx context.Context, customerID uint, data map[string]interface{}) error {
	db := infra.GetDB(ctx, r.db)

	subQuery := db.Model(&entity.Cart{}).
		Select("id").
		Where("customer_id = ?", customerID)

	result := db.Model(&entity.CartItem{}).
		Where("cart_id IN (?)", subQuery).
		Where("quantity > ?", 0).
		Updates(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
