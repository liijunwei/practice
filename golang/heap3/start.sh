#!/bin/bash

# 检查操作系统类型
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    open heap_visualizer.html
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    if command -v xdg-open &> /dev/null; then
        xdg-open heap_visualizer.html
    else
        echo "无法打开浏览器，请手动打开 heap_visualizer.html 文件"
    fi
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    # Windows
    start heap_visualizer.html
else
    echo "无法识别的操作系统，请手动打开 heap_visualizer.html 文件"
fi

echo "堆可视化工具已启动，如果浏览器没有自动打开，请手动打开 heap_visualizer.html 文件"
