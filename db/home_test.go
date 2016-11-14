package db

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetRollCommend(t *testing.T) {
	assert := assert.New(t)
	result, err := GetRollCommend()
	assert.Nil(err, "获取错误：%v", err)
	assert.Equal(8, len(result), "获取数据库数据错误")
	for i := 0; i < len(result); i++ {
		log.Infof("id:%v,url:%v,dest:%v", result[i].ID, result[i].ImageURL, result[i].Dest)
	}
}
