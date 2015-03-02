#!/bin/bash

release="RELEASE-1.2"
packages=( git supervisor python2.7 python-dev python-pip )

# Define the download command
if [ $(which curl) ]; then
    dl_cmd="curl -f"
else
    dl_cmd="wget --quiet"
fi

# Distribution detection
KNOWN_DISTRIBUTION="(Debian|Ubuntu|RedHat|CentOS|openSUSE|Amazon)"
DISTRIBUTION=$(lsb_release -d 2>/dev/null | grep -Eo $KNOWN_DISTRIBUTION  || grep -Eo $KNOWN_DISTRIBUTION /etc/issue 2>/dev/null || uname -s)
if [ -f /etc/debian_version -o "$DISTRIBUTION" == "Debian" -o "$DISTRIBUTION" == "Ubuntu" ]; then
    OS="Debian"
elif [ -f /etc/redhat-release -o "$DISTRIBUTION" == "RedHat" -o "$DISTRIBUTION" == "CentOS" -o "$DISTRIBUTION" == "openSUSE" -o "$DISTRIBUTION" == "Amazon" ]; then
    OS="RedHat"
elif [ -f /etc/system-release -o "$DISTRIBUTION" == "Amazon" ]; then
    # Some newer distros like Amazon may not have a redhat-release file
    OS="RedHat"
fi

# Define the install command
if [ $OS == "RedHat" ]; then
    $install_cmd="yum -y install"
    $install_ok="is not installed"
    $install_check_cmd="rpm -qi"
elif [$OS == "Debian"]; then
    $install_cmd="apt-get --force-yes --yes install"
    $install_ok="install ok installed"
    $install_check_cmd="dpkg-query -W --showformat='${Status}\n'"
fi

# Define elevation command (to root)
sudo_cmd=''
if [ $(echo "$UID") = "0" ]; then
    sudo_cmd='sudo'
fi

# Iterate over packages, check if they're installed, if not install
for i in "${packages[@]}"
do
    :
    PKG_OK=$(${sudo_cmd} ${install_check_cmd} ${i}|grep ${install_ok})
    if [ "" == "$PKG_OK" ]; then
        echo "Installing ${i}"
        ${sudo_cmd} ${install_cmd} ${i}
    else
        echo "Detected ${i} installed already"
    fi;
done

echo "Cloning git branch ${release}"
git clone https://bitbucket.org/sphire-development/serverstatusemitter.git -b ${release} --single-branch

echo "Copying SSE_Python_supervisord.conf to /etc/supervisor/conf.d"
cp SSE_Python_supervisord.conf /etc/supervisor/conf.d/SSE_Python_supervisord.conf

echo "Rereading and updating supervisorctl"
supervisorctl reread && supervisorctl update

# Install any requirements
echo "Installing project requirements via pip"
pip install -r requirements.txt
