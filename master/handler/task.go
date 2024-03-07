package handler

import (
	"context"
	"cron_tab_c/master/common"
	"cron_tab_c/master/global"
	"encoding/json"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
	etcd "go.etcd.io/etcd/client/v3"
	"net/http"
	"time"
)

func GetTaskList(ctx *gin.Context) {
	var (
		getResponse *etcd.GetResponse
		err         error
		data        []common.TaskVar
		tv          common.TaskVar
		total       int64
		e           *base.SentinelEntry
		b           *base.BlockError
	)
	if e, b = sentinel.Entry("[TASK_GET_LIST]", sentinel.WithTrafficType(base.Inbound)); b != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg":   "success",
				"tasks": "流量限制",
			},
		)
		return
	}

	data = make([]common.TaskVar, 0)
	if getResponse, err = global.KV.Get(context.Background(), "/Task/cron/", etcd.WithPrefix()); err == nil {
		total = getResponse.Count
		for _, v := range getResponse.Kvs {
			tv = common.TaskVar{}
			if err = json.Unmarshal(v.Value, &tv); err != nil {
				panic(err)
			}
			data = append(data, tv)
		}
	}
	e.Exit()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg":   "success",
			"tasks": data,
			"total": total,
		},
	)

}

func UpdateTask(ctx *gin.Context) {
	var (
		err         error
		putResponse *etcd.PutResponse
		taskReq     common.TaskRequest
		data        []byte
	)
	taskReq = common.TaskRequest{}
	taskReq.TaskVar.NextTime = time.Now()
	if err = ctx.ShouldBind(&taskReq); err != nil {
		fmt.Println(err)
		goto ClientERR
	}
	if data, err = json.Marshal(taskReq.TaskVar); err != nil {
		goto ERR
	}
	if putResponse, err = global.KV.Put(context.Background(), common.TaskSaveDir+taskReq.TaskVar.Name, string(data), etcd.WithPrevKV()); err != nil {
		goto ERR
	}
	if putResponse.PrevKv != nil {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg": fmt.Sprintf("update task: %v success", taskReq.TaskVar.Name),
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"msg": fmt.Sprintf("save task: %v success", taskReq.TaskVar.Name),
		},
	)
	return
ERR:
	ctx.JSON(
		http.StatusBadRequest,
		gin.H{
			"msg": "server err",
		},
	)
	return
ClientERR:
	ctx.JSON(
		http.StatusBadRequest,
		gin.H{
			"msg": "mission params",
		},
	)
	return
}

func DelJob(ctx *gin.Context) {
	var (
		tr          common.TaskRequest
		err         error
		delResponse *etcd.DeleteResponse
		total       int
	)
	tr = common.TaskRequest{}
	if err = ctx.ShouldBind(&tr); err != nil {
		goto ERR
	}

	if delResponse, err = global.KV.Delete(context.Background(), common.TaskSaveDir+tr.TaskVar.Name); err != nil {
		goto ERR
	}
	fmt.Println(delResponse.Deleted)
	if delResponse.Deleted != 0 {
		total = int(delResponse.Deleted)
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg":   "remove success",
				"total": total,
			},
		)
	} else {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"msg": "remove fail not have task",
			},
		)
	}
	return
ERR:
	ctx.JSON(
		http.StatusBadRequest,
		gin.H{
			"msg": "server err",
		},
	)
	return
}
