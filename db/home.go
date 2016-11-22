package db

import (
	"cainiaoshop-server/model"

	log "github.com/Sirupsen/logrus"
)

//GetRollRecommend 获取滚动推荐数据
func GetRollRecommend() ([]*model.RollRecommendModel, error) {
	err := DB.Ping()
	if err != nil {
		log.Debugf("GetRollRecommend 连接数据库错误：%v", err)
		return nil, err
	}

	var commends []*model.RollRecommendModel
	rows, err := DB.Query("SELECT * FROM t_roll_recommend")
	if err != nil {
		log.Debugf("getRollCommend 获取结果错误：%v", err)
		return nil, err
	}

	for rows.Next() {
		recommend := &model.RollRecommendModel{}
		err := rows.Scan(&recommend.ID, &recommend.ImageURL, &recommend.Dest)
		if err != nil {
			log.Infof("获取recommend错误：%v", err)
			return nil, err
		}
		commends = append(commends, recommend)

	}

	return commends, nil
}
