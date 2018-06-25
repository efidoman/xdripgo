#!/bin/bash

bt-device -r DexcomFE
systemctl restart bluetooth.service
sleep 10
echo
echo "Starting 1st run"
date
echo
./g5 410BFE
echo "Finished 1st run"
date
echo
echo
echo
bt-device -r DexcomFE
systemctl restart bluetooth.service
sleep 120

echo
echo "Starting 2nd run"
date
echo
./g5 410BFE
echo "Finished 2nd run"
date
echo
echo
echo
bt-device -r DexcomFE
systemctl restart bluetooth.service
sleep 120

echo
echo "Starting 3rd run"
date
echo
./g5 410BFE
echo "Finished 3rd run"
date
echo
echo
echo
#bt-device -r DexcomFE
#systemctl restart bluetooth.service
