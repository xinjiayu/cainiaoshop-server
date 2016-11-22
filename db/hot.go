package db

import (
	"cainiaoshop-server/model"

	log "github.com/Sirupsen/logrus"
)

//GetHotByLimit 通同限制获取更多数据
func GetHotByLimit(start int) ([]*model.HotModel, error) {
	err := DB.Ping()
	if err != nil {
		log.Debugf("GetHotByLimit 连接数据库错误：%v", err)
		return nil, err
	}
	var hots []*model.HotModel
    DB.Prepare("")
	return hots, nil
}
