package relay

import "gorm.io/gorm"

func getTotalCount[Model any](db *gorm.DB) (int64, error) {
	var totalCount int64
	var model Model

	newSession := db.Session(&gorm.Session{Initialized: true})
	if err := newSession.Model(&model).Count(&totalCount).Error; err != nil {
		return 0, err
	}

	return totalCount, nil
}
