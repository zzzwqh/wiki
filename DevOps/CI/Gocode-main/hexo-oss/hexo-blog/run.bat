@echo off
chcp 65001
echo 请选择选项： 1. 新建文章标题 2. 同步文件到远程仓库


set /p choice="请输入选项编号: "

if "%choice%"=="1" (
    set /p title="请输入文章标题: "
    hexo new "%title%"
) else if "%choice%"=="2" (
    set /p sync="是否同步本地内容到远程仓库？(y/n): "
    if /i "%sync%"=="y" (
        git add .
        git commit -m "blog changed"
        git push origin main
    ) else if /i "%sync%"=="n" (
        exit
    )
) else (
    echo 无效的选项，请重新运行脚本并输入 1 或者 2.
)
