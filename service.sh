#!/bin/bash

# 定义函数：复制配置文件并启动服务
copy_config_and_start() {
    local config_file_path="./manifest/config/config.yaml"

    # 删除现有的 config.yaml 文件
    rm -f "$config_file_path"

    # 根据参数决定复制哪个配置文件
    if [ "$2" == "online" ]; then
        echo "haha"
        cp "./manifest/config/config-online.yaml" "$config_file_path"
        # 替换 config.yaml 文件中的 {LOCAL_IP} 占位符为实际的 IP 地址
        sed -i "s/{LOCAL_IP}/$ip/g" "$config_file_path"
    else
        cp "./manifest/config/config-offline.yaml" "$config_file_path"
    fi

}

# 定义函数：启动 Go 应用程序
start_go_app() {
    go run main.go
}

start_go_app_online() {
  ./main
}

# 检查是否提供了启动命令
case "$1" in
    start)
        copy_config_and_start "$1"
        if [ $? -eq 0 ]; then
            start_go_app
        fi
        ;;
    startonline)
        copy_config_and_start "$1" online
        if [ $? -eq 0 ]; then
            start_go_app_online
        fi
        ;;
    restart)
        rm -f "./manifest/config/config.yaml"
        . "$0" start
        ;;
    restartonline)
        rm -f "./manifest/config/config.yaml"
        . "$0" startonline
        ;;
    *)
        echo "Usage: sh service.sh {start|startonline|restart|restartonline}"
        ;;
esac