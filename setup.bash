#!/usr/bin/bash

sudo apt update && sudo apt upgrade -y
sudo apt install git curl

cd ~
git clone https://github.com/yoshiyuki-140/isPrimeNumber-gRPC.git
cd isPrimeNumber

# setup python environment 
python -m venv venv
. ./venv/bin/activate
# install python dependencies
python -m pip intall -r requirements.txt

echo "Success installing python dependencies"

# install go dependencies
curl https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' > ~/.extraProfile
echo 'source ~/.extraProfile' > ~/.bashrc
. ~/.bashrc

# check go version
go version

# path check
echo $PATH | grep go

echo "Success!"
