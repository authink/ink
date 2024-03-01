package model

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/authink/inkstone"
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

type Log struct {
	inkstone.Model
	Detail string
}

func NewLog(log fmt.Stringer) *Log {
	return &Log{
		Detail: log.String(),
	}
}
