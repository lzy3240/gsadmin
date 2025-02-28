package baseservice

import (
	"github.com/jinzhu/gorm"
	"gsadmin/core/baseservice/condition"
	"reflect"
)

type Service struct {
}

// SetPaginate 设置分页
func (s *Service) SetPaginate(limit, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 设置默认值
		if limit == 0 {
			limit = 10
		}

		if page == 0 {
			page = 1
		}

		offset := (page - 1) * limit
		if offset < 0 {
			offset = 0
		}
		return db.Offset(offset).Limit(limit)
	}
}

// SetCondition 设置查询条件
func (s *Service) SetCondition(q interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		c := &condition.GormCondition{
			GormPublic: condition.GormPublic{},
			Join:       make([]*condition.GormJoin, 0),
		}
		condition.ResolveSearchQuery(q, c)
		for _, join := range c.Join {
			if join == nil {
				continue
			}
			db = db.Joins(join.JoinOn)
			for k, v := range join.Where {
				db = db.Where(k, v...)
			}
			for k, v := range join.Or {
				db = db.Or(k, v...)
			}
			for _, o := range join.Order {
				db = db.Order(o)
			}
		}
		for k, v := range c.Where {
			db = db.Where(k, v...)
		}
		for k, v := range c.Or {
			db = db.Or(k, v...)
		}
		for _, o := range c.Order {
			db = db.Order(o)
		}
		return db
	}
}

// 暂不需要
//func (s *Service) SetOrder() func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//
//		return db
//	}
//}

// gorm更新方法updates忽略零值, 故对象转换为map更新
func (s *Service) StructToMapByTag(obj interface{}, tagName string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(tagName) == "" {
			continue
		}
		data[t.Field(i).Tag.Get(tagName)] = v.Field(i).Interface()
	}
	return data
}
