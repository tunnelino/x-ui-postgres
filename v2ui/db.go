package v2ui

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"x-ui/config"
)

var v2db *gorm.DB

func initDB(dbPath string) error {
	c := &gorm.Config{
		Logger: logger.Discard,
	}
	var err error
	v2db, err = gorm.Open(postgres.Open(config.GetDBConn()), c)
	if err != nil {
		return err
	}

	return nil
}

func getV2Inbounds() ([]*V2Inbound, error) {
	inbounds := make([]*V2Inbound, 0)
	err := v2db.Model(V2Inbound{}).Find(&inbounds).Error
	return inbounds, err
}
