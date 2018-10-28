#!/bin/bash

sudo docker rmi grayDeploy/test-v1:latest
sudo docker build . -t grayDeploy/test-v1:latest
sudo docker push  grayDeploy/test-v1:latest

exit 0