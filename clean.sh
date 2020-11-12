#!/usr/bin/env bash

if [[ $1 != "" ]]
then
    [[ $1 = \/* ]] && echo "We don't work in / or below, aborting." && exit
    echo "Cleaning up $1"
    rm -rf $1
else
    echo "Cleaning up files"
    rm -rf index.html mission tunneling-comparison usage
fi
