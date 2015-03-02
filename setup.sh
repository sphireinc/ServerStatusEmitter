#!/bin/sh

packages=( git supervisor python2.7 python-dev python-pip )

# Iterate over packages, check if they're installed, if not install
for i in "${packages[@]}"
do
    :
    PKG_OK=$(dpkg-query -W --showformat='${Status}\n' ${i}|grep "install ok installed")
    if [ "" == "$PKG_OK" ]; then
        echo "Installing ${i}"
        apt-get --force-yes --yes install ${i}
    else
        echo "Detected ${i} installed already"
    fi;

    echo ""
done

echo "Cloning git branch RELEASE-1.1"
git clone https://bitbucket.org/sphire-development/serverstatusemitter.git -b RELEASE-1.1 --single-branch

echo "Copying SSE_Python_supervisord.conf to /etc/supervisor/conf.d"
cp SSE_Python_supervisord.conf /etc/supervisor/conf.d/SSE_Python_supervisord.conf

echo "Rereading and updating supervisorctl"
supervisorctl reread && supervisorctl update

# Install any requirements
echo "Installing project requirements via pip"
pip install -r requirements.txt
