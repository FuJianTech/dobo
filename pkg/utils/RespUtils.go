package utils

import (
	"net/http"
	"time"
)

// 错误相应模型
type APIResp struct {
	StatusCode int64        `json:"status_code"`
	Msg  string				`json:"msg"`
	Data interface{}		`json:"data"`
	Timestamp int64			`json:"timestamp"`

}

func RespError(statsCode int64, err error) APIResp {
	return APIResp{StatusCode: statsCode, Msg: err.Error(), Timestamp: time.Now().Unix() }
}

func RespErrorMsg(statsCode int64, err string) APIResp {
	return APIResp{StatusCode: statsCode, Msg: err, Timestamp: time.Now().Unix()}
}

func RespData(statsCode int64, data interface{}) APIResp {
	return APIResp{StatusCode: statsCode , Data: data, Timestamp: time.Now().Unix()}
}

func RespSuccess() APIResp {
	return APIResp{StatusCode: http.StatusOK, Timestamp: time.Now().Unix()}
}
