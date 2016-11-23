package db

import (
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetHotByLimit(t *testing.T) {
	assert := assert.New(t)
	r, err := GetHotByLimit(0)
	assert.Nil(err, "Test GetHotByLimit 出现错误")
	for _, v := range r {
		log.Infof("v:%v ", v)
	}
}
