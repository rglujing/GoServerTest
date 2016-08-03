#!/bin/bash

Ret=`grep result AuCode.txt | awk '{for (i=3;i<=NF;i++) {printf $i}}'`
aucode=`echo  "$Ret" | ./JSON.sh -l | grep AuCode | awk '{print $2}' | tr -d '"'`

echo "$aucode"

msg=`InterGration -m RS -aucode $aucode $1`
echo "RequestSession $msg"

Err=`echo $msg | grep err`
if test ! "$Err"X = X; then
	echo "Error happened"
	exit
fi
echo $msg > AsToken.txt
Ret=`grep result AsToken.txt | awk '{for (i=3;i<=NF;i++) {printf $i}}'`
astoken=`echo  "$Ret" | ./JSON.sh -l | grep AsToken | awk '{print $2}' | tr -d '"'`

echo "astoken is " $astoken
