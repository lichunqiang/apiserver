#!/usr/bin/env bash

BIN="apiserver"
DIR=$PWD
INTERVAL=2

ARGS=""

function start()
{
    if [ "`pgrep $BIN -u $UID`" != "" ]; then
        echo "$BIN is already running"
        exit 1
    fi

    nohup $DIR/$BIN $ARGS server &>/dev/null &

    echo "Staring..." && sleep $INTERVAL

    # check status
	if [ "`pgrep $BIN -u $UID`" == "" ];then
		echo "$BIN start failed"
		exit 1
	fi

	echo "$BIN started successfully."
}

function status()
{
    if [ "`pgrep $BIN -u $UID`" != "" ]; then
        echo $BIN is running
    else
        echo $BIN is stopped
    fi
}

function stop()
{
    if [ "`pgrep $BIN -u $UID`" != "" ];then
		kill -9 `pgrep $BIN -u $UID`
	fi

	echo "stopping..." &&  sleep $INTERVAL

	if [ "`pgrep $BIN -u $UID`" != "" ];then
		echo "$BIN stop failed"
		exit 1
	fi
}

function help()
{
    echo "usage: $1 {start|stop|restart|status}"
	exit 1
}

case "$1" in
    start)
        start
    ;;
    status)
        status
    ;;
    stop)
      stop
    ;;
    restart)
        stop && start
    ;;
    *)
        help $0
    ;;
esac