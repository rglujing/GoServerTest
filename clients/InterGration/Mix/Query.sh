#!/bin/bash

Ret=`grep result AsToken.txt | awk '{for (i=3;i<=NF;i++) {printf $i}}'`
astoken=`echo  "$Ret" | ./JSON.sh -l | grep AsToken | awk '{print $2}' | tr -d '"'`

echo "astoken is " $astoken

msg=`InterGration -m AccessToken  -astoken $astoken $1`
echo "AccessToken  $msg"
msg=`InterGration -m InspectToken -astoken $astoken $1`
echo "InspectToken $msg"


