package dao

import (
	"errors"

	"cnb.cool/mliev/open/dwz-server/app/model"
	"cnb.cool/mliev/open/dwz-server/internal/interfaces"
)

// DomainDao 域名DAO
type DomainDao struct {
	helper interfaces.HelperInterface
}

func NewDomainDao(helper interfaces.HelperInterface) *DomainDao {
	return &DomainDao{helper: helper}
}

// Create 创建域名
func (d *DomainDao) Create(domain *model.Domain) error {
	return d.helper.GetDatabase().Create(domain).Error
}

// FindByDomain 根据域名查找
func (d *DomainDao) FindByDomain(domain string) (*model.Domain, error) {
	var domainModel model.Domain
	err := d.helper.GetDatabase().Where("domain = ? AND deleted_at IS NULL", domain).First(&domainModel).Error
	if err != nil {
		return nil, err
	}
	return &domainModel, nil
}

// List 获取域名列表
func (d *DomainDao) List() ([]model.Domain, error) {
	var domains []model.Domain
	err := d.helper.GetDatabase().Where("deleted_at IS NULL").Order("created_at DESC").Find(&domains).Error
	return domains, err
}

// Update 更新域名
func (d *DomainDao) Update(domain *model.Domain) error {
	return d.helper.GetDatabase().Save(domain).Error
}

// Delete 删除域名（软删除）
func (d *DomainDao) Delete(id uint64) error {
	return d.helper.GetDatabase().Delete(&model.Domain{}, id).Error
}

// GetActiveDomains 获取活跃域名列表
func (d *DomainDao) GetActiveDomains() ([]model.Domain, error) {
	var domains []model.Domain
	database := d.helper.GetDatabase()
	if database != nil {
		return nil, errors.New("database Connect Error")
	}
	err := database.Where("is_active = ? AND deleted_at IS NULL", true).Find(&domains).Error
	return domains, err
}

// ExistsByDomain 检查域名是否已存在
func (d *DomainDao) ExistsByDomain(domain string) (bool, error) {
	var count int64
	err := d.helper.GetDatabase().Model(&model.Domain{}).
		Where("domain = ? AND deleted_at IS NULL", domain).Count(&count).Error
	return count > 0, err
}
