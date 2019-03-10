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

echo -e "
#=================================================
#       System Required: CentOS/Debian/Ubuntu/Darwin
#       Description: AUXPI build
#       Version: 1.0.0
#       Author: aimerforreimu
#       Blog: https://0w0.tn
#=================================================
"

function buildHelp(){
    echo "Auxpi Build Helper V1.0"
    echo
    echo "all [version] [clear] ---- Build all platforms programs"
    echo "tar [version] ---- Tar all platforms programs"
    echo "mac ---- Build mac program"
    echo "mac ---- Build mac program"
    echo "mac ---- Build mac program"
    echo "linux ---- Build windows program"
    echo "windows ---- Build linux program"
    echo "clear ---- Delete 'build/' folder"
    echo "help ---- Show help info"
    echo
}

function buildAndMove() {
echo -e "${Info_font_prefix}[INFO:]Begin to compile ${1} program ${Font_suffix} "
remove $1
echo -e "${Info_font_prefix}[INFO:]Clear File Done  ${Font_suffix} "
GOOS=$1 GOARCH=amd64 go  build main.go
echo -e "${Info_font_prefix}[INFO:]Build ${1} program done ${Font_suffix} "
mkdir -p build/$1
mkdir build/$1/conf
echo -e "${Info_font_prefix}[INFO:]Create folder done ${Font_suffix} "
if [ "$1"x = "windows"x ] ;then
    mv main.exe build/$1/auxpi.exe
    else
    mv main build/$1/auxpi
fi

cp -r static/ build/$1/static/
cp -r views/ build/$1/views/
cp -r conf/app.conf build/$1/conf/

cp LICENSE build/$1/
cp README.MD build/$1/

echo -e "${Info_font_prefix}[INFO:]Copy file done ${Font_suffix} "
echo -e "${Info_background_prefix}[INFO:]Done all work! : ) ${Font_suffix} "
echo -e "#======================================================#"

echo -e "${Info_font_prefix}[SUCCESS:]Your can see your ${1} program in 'build/${1}' ${Font_suffix} "

}

function remove() {
rm -rf build/$1
rm -rf auxpi
rm -rf main
rm -rf main.exe
}

function clearAll() {
    rm -rf build/
    echo -e "${Info_font_prefix}[INFO:]Delete 'build/' folder ${Font_suffix} "
}

function tarAll() {
    tar zcvf build/auxpi-${1}-darwin.tar.gz build/darwin
    echo -e "${Info_font_prefix}[INFO:]Tar darwin done ${Font_suffix} "
    tar zcvf build/auxpi-${1}-linux.tar.gz build/linux
    echo -e "${Info_font_prefix}[INFO:]Tar linux done ${Font_suffix} "
    tar zcvf build/auxpi-${1}-windows.tar.gz build/windows
    echo -e "${Info_font_prefix}[INFO:]Tar windows done ${Font_suffix} "
}
function buildAll() {
     buildAndMove darwin
     buildAndMove linux
     buildAndMove windows
     tarAll $1
     if [ "$2"x = "clear"x ]; then
         buildClear
     fi
}
function buildClear(){
    rm -rf build/darwin
    rm -rf build/linux
    rm -rf build/windows
    echo -e "${Info_font_prefix}[INFO:]Clear build files done ${Font_suffix} "
}

function buildAdmin() {
    cd resource
    yarn run build
    cd ..
    echo -e "${Info_font_prefix}[INFO:]Admin Build Done!${Font_suffix} "
}

# Initialization step
action=$1
version=$2
clear=$3

[ -z $1 ] && action=linux
case "$action" in
mac)
    buildAndMove darwin
    ;;
windows)
    buildAndMove windows
    ;;
linux)
    buildAndMove linux
    ;;
clear)
    clearAll
    ;;
help)
    buildHelp
    ;;
tar)
    tarAll $2
    ;;
all)
    buildAll $2 $3
    ;;
admin)
    buildAdmin
    ;;
    *)
    echo -e "${Error_font_prefix}[INFO:]Parameter error , please use help to see how to use ${Font_suffix}"
    ;;
esac





