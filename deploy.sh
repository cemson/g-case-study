#!/bin/bash

remote_user="root"
remote_server="ec2-3-83-238-113.compute-1.amazonaws.com"
remote_dir="/home/ubuntu"

echo "Running unit tests"
go test -v ./...

echo "Packing source and compress"
mkdir g-case-study
cp -r clients g-case-study
cp -r consts g-case-study
cp -r controllers g-case-study
cp -r dto g-case-study
cp -r exceptions g-case-study
cp -r globals g-case-study
cp -r inMemoryStore g-case-study
cp -r logging g-case-study
cp -r repo g-case-study
cp -r services g-case-study
cp -r settings g-case-study
cp -r utilities g-case-study
cp -r validators g-case-study
cp conf.prod.json g-case-study
cp go.mod g-case-study
cp go.sum g-case-study
cp main.go g-case-study
tar cfz g-case-study.tar.gz g-case-study

echo "Upload the app to the remote server"
scp g-case-study.tar.gz "${remote_user}@${remote_server}:${remote_dir}"

echo "SSH into the remote server and restart the systemd service"
ssh "${remote_user}@${remote_server}" "rm -rf ${remote_dir}/g-case-study"
ssh "${remote_user}@${remote_server}" "tar xfz ${remote_dir}/g-case-study.tar.gz -C ${remote_dir}"
ssh "${remote_user}@${remote_server}" "rm ${remote_dir}/g-case-study.tar.gz"
ssh "${remote_user}@${remote_server}" "export PATH=\$PATH:/usr/local/go/bin && cd ${remote_dir}/g-case-study && go build"
ssh "${remote_user}@${remote_server}" "systemctl restart getir-api.service"

rm -rf g-case-study
rm g-case-study.tar.gz

echo "Deployment completed."
