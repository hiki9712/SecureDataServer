package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

var ctx = gctx.New()

type data struct {
	index int
	value string
}

var dataList map[int]*data

// 模拟往内存中加有了数据
func store(ctx context.Context) {
	dataList = make(map[int]*data)
	data1 := &data{
		index: 1,
		value: "lhy",
	}
	dataList[data1.index] = data1
	data2 := &data{
		index: 2,
		value: "lhy2",
	}
	dataList[data2.index] = data2
	g.Log().Info(ctx, "add data success!")
}

func main() {
	s := g.Server()
	s.BindHandler("/ws", func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(ctx, err)
			r.Exit()
		}
		store(ctx)
		for {
			msgType, msg, err := ws.ReadMessage()
			g.Log().Info(ctx, msgType, msg)
			if err != nil {
				return
			}
			if err = ws.WriteMessage(msgType, []byte(dataList[gconv.Int(string(msg))].value)); err != nil {
				return
			}
			delete(dataList, gconv.Int(string(msg)))
		}
	})
	s.SetServerRoot(gfile.MainPkgPath())
	s.SetPort(8199)
	s.Run()
}
