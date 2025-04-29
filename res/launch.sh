#!/bin/bash

pkill -f hwt

if [ "$#" -gt 0 ]; then
 /mnt/SDCARD/Apps/Rubik/rubik "$@"
else
  progdir=$(dirname "$0")
  cd $progdir
  ./rubik
fi