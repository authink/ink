package models

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/authink/orm/model"
	"github.com/gin-gonic/gin"
)

type LogDetail struct {
	AppId     int        `json:"appId,omitempty"`
	StaffId   int        `json:"staffId,omitempty"`
	Resource  string     `json:"resource,omitempty"`
	Action    string     `json:"action,omitempty"`
	PathVars  gin.Params `json:"pathVars,omitempty"`
	QueryVars url.Values `json:"queryVars,omitempty"`
}

func (ld *LogDetail) String() string {
	bytes, err := json.Marshal(ld)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// @model
// @db s_logs
type Log struct {
	model.Created
	Detail string
}

func NewLog(log fmt.Stringer) *Log {
	return &Log{
		Detail: log.String(),
	}
}

func (log *Log) GetDetail() *LogDetail {
	detail := &LogDetail{}
	json.Unmarshal([]byte(log.Detail), detail)
	return detail
}
