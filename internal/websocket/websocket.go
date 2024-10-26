// internal/websocket/websocket.go
package websocket

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	wsConnections = make(map[int64]*ghttp.WebSocket)
	wsMutex       sync.Mutex
)

// WebSocket 处理函数
func HandleWebSocket(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		g.Log().Error(r.Context(), "WebSocket connection failed:", err)
		r.Exit()
	}

	// // 假设 TaskID 是连接的唯一标识符
	// taskID := gconv.Int64(r.Get("TaskID"))
	// if taskID == 0 {
	// 	g.Log().Error(r.Context(), "TaskID is required")
	// 	r.Exit()
	// }

	wsMutex.Lock()
	wsConnections[0] = ws
	wsMutex.Unlock()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			g.Log().Error(r.Context(), "WebSocket read error:", err)
			break
		}
		g.Log().Info(r.Context(), "Received message:", string(msg))

		// 处理消息并发送响应
		if err = ws.WriteMessage(msgType, msg); err != nil {
			g.Log().Error(r.Context(), "WebSocket write error:", err)
			break
		}
	}

	wsMutex.Lock()
	delete(wsConnections, 0)
	wsMutex.Unlock()
}

// 发送 WebSocket 消息
func SendWebSocketMessage(ctx context.Context, taskID int64, resultValue string) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	// g.Log().Info(ctx, "111111111")

	if ws, ok := wsConnections[0]; ok {
		// 将 TaskID 和 resultValue 打包成 JSON 对象
		g.Log().Info(ctx, "connection found")
		message := map[string]interface{}{
			"TaskID":      taskID,
			"ResultValue": resultValue,
		}
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			g.Log().Error(ctx, "Failed to marshal message:", err)
			return
		}

		if err := ws.WriteMessage(ghttp.WsMsgText, jsonMessage); err != nil {
			g.Log().Error(ctx, "WebSocket write error:", err)
		}
	} else {
		g.Log().Warning(ctx, "WebSocket connection not found for TaskID:", taskID)
	}
}
