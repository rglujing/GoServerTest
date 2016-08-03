#!/bin/bash

msg=`InterGration -m C $1`
echo "CreateAuCode $msg"

Err=`echo $msg | grep err`
if test ! "$Err"X = X; then
	echo "Error happened"
	exit
fi

echo $msg > AuCode.txt

Ret=`grep result AuCode.txt | awk '{for (i=3;i<=NF;i++) {printf $i}}'`
aucode=`echo  "$Ret" | ./JSON.sh -l | grep AuCode | awk '{print $2}' | tr -d '"'`
echo "aucode is $aucode"
