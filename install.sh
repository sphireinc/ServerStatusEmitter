#!/usr/bin/bash

if [ "$(whoami)" != "root" ]; then
	echo "You must run this as root."
	exit 1
fi

apt-get install python2.7
apt-get install python-pip
pip install psutil
chmod a-x ./install.sh