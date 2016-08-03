#!/bin/bash

Ret=`grep result AsToken.txt | awk '{for (i=3;i<=NF;i++) {printf $i}}'`
astoken=`echo  "$Ret" | ./JSON.sh -l | grep AsToken | awk '{print $2}' | tr -d '"'`

echo "astoken is " $astoken

msg=`InterGration -m AccessToken  -astoken $astoken`
echo "AccessToken  $msg"
msg=`InterGration -m InspectToken -astoken $astoken`
echo "InspectToken $msg"


