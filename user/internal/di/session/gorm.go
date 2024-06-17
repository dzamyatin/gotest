package session

import (
	"app/user/internal/di/static"
	"gorm.io/gorm"
)

func (s *Session) NewGormSession() *gorm.DB {
	return getOrCreateTyped(s, func() *gorm.DB {
		return static.GetGorm().WithContext(s.ctx)
	})
}
