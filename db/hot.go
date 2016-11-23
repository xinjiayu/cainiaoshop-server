package db

import (
	"cainiaoshop-server/model"

	log "github.com/Sirupsen/logrus"
)

//GetHotByLimit 通同限制获取更多数据
func GetHotByLimit(page int) ([]*model.HotModel, error) {
	err := DB.Ping()
	if err != nil {
		log.Debugf("GetHotByLimit 连接数据库错误：%v", err)
		return nil, err
	}
	var hots []*model.HotModel
	if page <= 0 {
		page = 0
	}
	rows, err := DB.Query("SELECT * FROM cainiaoshop.t_hot_model limit 10 offset ?", page*10)
	if err != nil {
		log.Debugf("GetHotByLimit 生成语句错误：%v", err)
		return nil, err
	}

	for rows.Next() {
		hot := &model.HotModel{}
		err := rows.Scan(&hot.ID, &hot.Price, &hot.PicURL, &hot.Dest)
		if err != nil {
			log.Debugf("GetHotByLimit 获取数据错误：%v", err)
			continue
		}
		hots = append(hots, hot)
	}
	return hots, nil
}
