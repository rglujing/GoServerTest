#!/bin/bash

./InterGration -m C > AuCode.txt

Err=`grep err AuCode.txt`
if test ! "$Err"X = X; then
	echo "Error happened"
	exit
fi

Ret=`grep result AuCode.txt | awk '{for (i=3;i<=NF;i++) {printf $i}}'`
aucode=`echo  "$Ret" | ./JSON.sh -l | grep AuCode | awk '{print $2}' | tr -d '"'`

echo "$aucode"

./InterGration -m RT -aucode $aucode > AsToken.txt
Err=`grep err AuCode.txt`
if test ! "$Err"X = X; then
	echo "Error happened"
	exit
fi

Ret=`grep result AsToken.txt | awk '{for (i=3;i<=NF;i++) {printf $i}}'`
astoken=`echo  "$Ret" | ./JSON.sh -l | grep AsToken | awk '{print $2}' | tr -d '"'`

echo "astoken is " $astoken

./InterGration -m RefreshToken -expire=1000 -astoken $astoken
./InterGration -m AccessToken  -astoken $astoken
./InterGration -m InspectToken -astoken $astoken


