#!/bin/bash


#Git

#sudo apt-get update
#sudo apt-get install -y make
#sudo apt-get install -y git
#mkdir bin pkg
#mkdir -p src/github.com/hpcloud
#mv ~/id_rsa* ~/.ssh/
#cd src/github.com/hpcloud
#git clone git@github.com:hpcloud/hdp-resource-manager.git


#Go
#wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
#gunzip go1.6.2.linux-amd64.tar.gz
#tar -xvf go1.6.2.linux-amd64.tar
#sudo mv go /usr/local/
#cat <<EOT >> ~/.profile
#export WORKDIR=/home/ubuntu/src/github.com/hpcloud/hdp-resource-manager
#export PATH=$PATH:/usr/local/go/bin
#export GOPATH=/home/ubuntu
#EOT

#source ~/.profile


#Docker
#sudo apt-get install apt-transport-https ca-certificates
#sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
#sudo bash -c 'cat <<EOT >> /etc/apt/sources.list.d/docker.list
#deb https://apt.dockerproject.org/repo ubuntu-trusty main
#EOT'

#sudo apt-get update
#sudo apt-get install linux-image-extra-$(uname -r)
#sudo apt-get install docker-engine
#sudo service docker start
#sudo groupadd docker
#sudo usermod -aG docker ubuntu



#swagger
#sudo apt-get install -y npm
#sudo npm install -g swagger-cli


#Zip
#sudo apt-get install -y zip


#Vim with lua
sudo apt-get remove --purge vim vim-runtime vim-gnome vim-tiny vim-common vim-gui-common
sudo apt-get install -y liblua5.1-dev luajit libluajit-5.1 python-dev ruby-dev libperl-dev mercurial libncurses5-dev libgnome2-dev libgnomeui-dev libgtk2.0-dev libatk1.0-dev libbonoboui2-dev libcairo2-dev libx11-dev libxpm-dev libxt-dev
sudo mkdir /usr/include/lua5.1/include
sudo ln -s /usr/include/luajit-2.0 /usr/include/lua5.1/include
cd ~
git clone https://github.com/vim/vim.git
cd vim/src

make distclean
./configure --with-features=huge \
           --enable-rubyinterp \
           --enable-largefile \
           --disable-netbeans \
           --enable-pythoninterp \
           --with-python-config-dir=/usr/lib/python2.7/config-x86_64-linux-gnu \
           --enable-perlinterp \
           --enable-luainterp \
           --with-luajit \
           --enable-gui=auto \
           --enable-fail-if-missing \
           --with-lua-prefix=/usr/include/lua5.1 \
           --enable-cscope
make
sudo make install
cd ..
sudo mkdir /usr/share/vim
sudo mkdir /usr/share/vim/vim74
sudo cp -fr runtime/* /usr/share/vim/vim74/
curl vimfiles.luan.sh/install | bash



#sudo reboot
