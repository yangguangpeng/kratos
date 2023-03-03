#!/bin/bash
Usage() {
    echo "Usage:"
    echo "start.sh [-e ENV]"
    echo "Description:"
    echo "ENV,运行环境,如：dev、fat、pre、pro"
    exit -1
}

while getopts :h:e: OPT; do
    case $OPT in
        e) env="$OPTARG";;
        h) Usage;;
        ?) Usage;;
    esac
done

if [ ! -n "$env" ]
then
  env=dev
fi

cp -f "./configs/"$env"/config.yaml" "configs/config.yaml"
echo "=====================环境:"$env"========================="
kratos run