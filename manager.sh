#!/bin/bash

SERVER="auxpi"
BASE_DIR=$PWD
INSTALL_DIR="/root/auxpi"
INTERVAL=2

ARGS=""

#color
INFO_FONT_PREFIX="\033[32m"
ERROR_FONT_PREFIX="\033[31m"
INFO_BACKGROUND_PREFIX="\033[42;37m"
ERROR_BACKGROUND_PREFIX="\033[41;37m"
FONT_SUFFIX="\033[0m"


#text block
INFO_BLOCK=${INFO_FONT_PREFIX}"[INFO]:"${FONT_SUFFIX}
ERROR_BLOCK=${ERROR_FONT_PREFIX}"[ERROR]:"${FONT_SUFFIX}

function start()
{
    if [[ "`pgrep $SERVER -u $UID`" != "" ]];then
		echo -e ${ERROR_BLOCK} "$SERVER already running"
		exit 1
    fi
    nohup ${BASE_DIR}/${SERVER} &> ${INSTALL_DIR}/auxpi.out &

    echo -e ${INFO_BLOCK} "sleeping & checking ..." && sleep ${INTERVAL}

    # check status
	if [[ "`pgrep $SERVER -u $UID`" == "" ]];then
		echo -e ${ERROR_BLOCK} "$SERVER start failed"
		echo -e ${INFO_BLOCK} "start by install dir"
		nohup ${INSTALL_DIR}/${SERVER} &> ${INSTALL_DIR}/auxpi.out &
		exit 1
	else
	    echo -e ${INFO_BLOCK} "start success"
	    exit 1
    fi

    echo -e ${INFO_BLOCK} "sleeping & checking ......" && sleep ${INTERVAL}

    # check status
	if [[ "`pgrep $SERVER -u $UID`" == "" ]];then
		echo -e ${ERROR_BLOCK} "$SERVER start failed"
		exit 1
	else
	   echo -e ${INFO_BLOCK} "start success"
    fi

}

function status()
{
	if [[ "`pgrep $SERVER -u $UID`" != "" ]];then
		echo -e ${ERROR_BLOCK} ${SERVER} is running
	else
		echo -e ${INFO_BLOCK} ${SERVER} is not running
	fi
}

function stop()
{


	if [[ "`pgrep $SERVER -u $UID`" != "" ]];then
		kill -9 `pgrep ${SERVER} -u $UID`
	else
		echo -e ${ERROR_BLOCK} ${SERVER} has already stopped.
		exit 1
	fi

	echo -e ${INFO_BLOCK} "sleeping & checking ......" &&  sleep ${INTERVAL}

	if [[ "`pgrep $SERVER -u $UID`" != "" ]];then
		echo -e ${ERROR_BLOCK} "$SERVER stop failed"
		exit 1
		else
		echo -e ${INFO_BLOCK} "stop success"
	fi
}

case "$1" in
	'start')
	start
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	stop && start
	;;
	*)
	echo "usage: $0 {start|stop|restart|status}"
	exit 1
	;;
esac
