package sync

import (
	"bytes"
	"encoding/json"
	"goclip/model"
	"goclip/pkg"
	"goclip/storage"
	"io"
	"net/http"
)

func SyncHistory() {

}

func SendRemote(url string) error {
	var historyList []model.History
	historyList = storage.GetData()
	//转换为json
	jsonBytes, _ := json.Marshal(historyList)
	//将数据发送到指定url
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
	}
	defer resp.Body.Close()
	return nil
}

func handleSync(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "请求方法不是POST", http.StatusMethodNotAllowed)
		return
	}

	body, _ := io.ReadAll(r.Body)
	var historyList []model.History
	//将数据重新转换为结构体
	err := json.Unmarshal(body, &historyList)
	if err != nil {
	}

	//ToDo 写入数据库
	for _, history := range historyList {
		storage.UpdateData(history.Id, history.Content, history.Time)
	}

	//写入完毕
	pkg.WriteInfoLog("service.log", "同步完成")
}

//ToDo 使用全局变量来区分到底是cli还是gui因为很多服务需要区分是在命令行来展示信息还是在gui中显示
