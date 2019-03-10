#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH
#=================================================
#       System Required: CentOS/Debian/Ubuntu/Darwin
#       Description: AUXPI build
#       Version: 1.0.0
#       Author: aimer
#       Blog: https://0w0.tn
#=================================================
Info_font_prefix="\033[32m" && Error_font_prefix="\033[31m" && Info_background_prefix="\033[42;37m" && Error_background_prefix="\033[41;37m" && Font_suffix="\033[0m"
install_path='/root/auxpi'
auxpi_path="${install_path}/build/linux"
name="auxpi"

echo -e "
#=================================================
#       System Required: CentOS/Debian/Ubuntu/Darwin
#       Description: AUXPI Install
#       Version: 1.0.0
#       Author: aimerforreimu
#       Blog: https://0w0.tn
#=================================================
"
function auxpi_help(){
    echo "Auxpi Installer V1.0"
    echo
    echo "all  ---- Install Nginx Mysql Auxpi"
    echo "install ---- Install Auxpi"
    echo "mysql ---- Install Mysql"
    echo "nginx ---- Install Nginx"
    echo "help ---- Show help info"
    echo
}

check_sys(){
	if [[ -f /etc/redhat-release ]]; then
		release="centos"
	elif cat /etc/issue | grep -q -E -i "debian"; then
		release="debian"
	elif cat /etc/issue | grep -q -E -i "ubuntu"; then
		release="ubuntu"
	elif cat /etc/issue | grep -q -E -i "centos|red hat|redhat"; then
		release="centos"
	elif cat /proc/version | grep -q -E -i "debian"; then
		release="debian"
	elif cat /proc/version | grep -q -E -i "ubuntu"; then
		release="ubuntu"
	elif cat /proc/version | grep -q -E -i "centos|red hat|redhat"; then
		release="centos"
    fi
	bit=$(uname -m)
}

check_root(){
	[[ $EUID != 0 ]] && echo -e "${Error} 当前非ROOT账号(或没有ROOT权限)，无法继续操作，请更换ROOT账号或使用 ${Green_background_prefix}sudo su${Font_color_suffix} 命令获取临时ROOT权限（执行后可能会提示输入当前账号的密码）。" && exit 1
}

function install() {
    [[  -e "${install_path}/auxpi" ]] && echo -e "${Error_font_prefix}[ERROR]${Font_suffix}已经安装无需再次安装" && exit 1
    check_root
    mkdir -p $install_path
    cd $install_path
    wget --no-check-certificate -O "auxpi.tar.gz" "https://github.com/aimerforreimu/AUXPI/releases/download/2.0.0/auxpi-2.0.0-linux.tar.gz"
    [[ ! -e "auxpi.tar.gz" ]] && echo -e "${Error_font_prefix}[ERROR]${Font_suffix} auxpi 下载失败" && exit 1
    tar zxvf "auxpi.tar.gz"
    rm -rf "auxpi.tar.gz"
    [[ ! -e ${auxpi_path} ]] && echo -e "${Error_font_prefix}[ERROR]${Font_suffix} auxpi 解压失败或压缩文件错误 !" && exit 1
    cd /root
    cp -r ${auxpi_path}/* ${install_path}
    [[ ! -e "${install_path}/auxpi" ]] && echo -e "${Error_font_prefix}[ERROR]${Font_suffix} auxpi 文件移动出错" && exit 1
    rm -rf "${install_path}/build"
    cd $install_path
    chmod +x auxpi
    ./auxpi init
    echo -e "${Info_font_prefix}[INFO]${Font_suffix} auxpi 初始化完成:"
    echo -e "${Info_font_prefix}安装路径: ${install_path}${Font_suffix}"
    echo -e "${Info_font_prefix}配置文件: ${install_path}/conf/siteConfig.json ${Font_suffix}"
    echo -e "${Info_font_prefix}其它教程: https://github.com/aimerforreimu/AUXPI/wiki ${Font_suffix}"
}

function install_mysql() {
    wget -c http://mirrors.linuxeye.com/oneinstack-full.tar.gz && tar xzf oneinstack-full.tar.gz && ./oneinstack/install.sh --db_option 3 --dbinstallmethod 1 --dbrootpwd 1s16r74z}

function install_nginx() {
    wget -c http://mirrors.linuxeye.com/oneinstack-full.tar.gz && tar xzf oneinstack-full.tar.gz && ./oneinstack/install.sh --nginx_option 1
}

function install_all() {
    check_root
    wget -c http://mirrors.linuxeye.com/oneinstack-full.tar.gz && tar xzf oneinstack-full.tar.gz && ./oneinstack/install.sh --nginx_option 1 --db_option 3 --dbinstallmethod 1 --dbrootpwd auxpi_password --redis  --memcached  --iptables
    install
    echo -e "${Info_font_prefix}数据库密码:auxpi_password    请记得及时更改您的数据库密码${Font_suffix}"
}

# Initialization step
check_sys
action=$1
[ -z $1 ] && action=help
case "$action" in
install)
    install
    ;;
    nginx)
    install_nginx
    ;;
    mysql)
    install_mysql
    ;;
    all)
    install_all
    ;;
help)
    auxpi_help
    ;;
    *)
    echo "用法错误! 用法请查看 help 。"
    ;;
esac