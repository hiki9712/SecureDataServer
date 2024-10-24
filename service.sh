#!/bin/bash

# 定义函数：复制配置文件并启动服务
copy_config_and_start() {
    local config_file_path="./manifest/config/config.yaml"
    local local_ip

    # 删除现有的 config.yaml 文件
    rm -f "$config_file_path"

    # 根据参数决定复制哪个配置文件
    if [ "$2" == "online" ]; then
        cp "./manifest/config/config-online.yaml" "$config_file_path"
    else
        cp "./manifest/config/config-offline.yaml" "$config_file_path"
    fi

    # 获取本机的 IP 地址
    if command -v ip >/dev/null 2>&1; then
        # 对于支持 ip 命令的系统
        local_ip=$(ip addr show eno1 | grep "inet\b" | awk '{print $2}' | cut -d/ -f1)
        if [ -z "$local_ip" ]; then
            local_ip=$(ip addr show en0 | grep "inet\b" | awk '{print $2}' | cut -d/ -f1)
        fi
    else
        # 对于 macOS，使用 ifconfig 命令
        local_ip=$(ifconfig eno1 | grep "inet\b" | awk '{print $2}' | cut -d: -f2)
        if [ -z "$local_ip" ]; then
            local_ip=$(ifconfig en0 | grep "inet\b" | awk '{print $2}' | cut -d: -f2)
        fi
    fi

    # 检查是否成功获取到 IP 地址
    if [ -z "$local_ip" ]; then
        echo "Error: Unable to find the local IP address."
        exit 1
    else
        echo "Local IP address is: $local_ip"
    fi

    # 替换 config.yaml 文件中的 {LOCAL_IP} 占位符为实际的 IP 地址
    sed -i "" "s/{LOCAL_IP}/$local_ip/g" "$config_file_path"

    echo "Updated $config_file_path with local IP address."
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