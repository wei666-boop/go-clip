package internal

import (
	"github.com/spf13/viper"
	"goclip/storage"
)

func CheckHistoryCount() {
	viper.SetConfigName("./../conf/clipboard.yaml")
	err := viper.ReadInConfig()
	if err != nil {

	}
	var maxCount int64 = viper.GetInt64("maxHistory")
	var nowCount int64
	if maxCount <= 0 {

	}
	//ToDO将数据库中超出部分删除
	nowCount, _ = storage.GetCount()
	if nowCount > maxCount {
		storage.DeleteData(nowCount - maxCount)
	}
}
