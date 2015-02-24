# Install python2.7 if not installed
PKG_OK=$(dpkg-query -W --showformat='${Status}\n' python2.7|grep "install ok installed")
if [ "" == "$PKG_OK" ]; then
    echo "Installing python2.7"
    apt-get --force-yes --yes install python2.7
else
    echo "Detected python2.7 installed already"
fi;

echo ""

# Install python-dev if not installed
PKG_OK=$(dpkg-query -W --showformat='${Status}\n' python-dev|grep "install ok installed")
if [ "" == "$PKG_OK" ]; then
    echo "Installing python-dev"
    apt-get --force-yes --yes install python-dev
else
    echo "Detected python-dev installed already"
fi;

echo ""

# Install python-pip if not installed
PKG_OK=$(dpkg-query -W --showformat='${Status}\n' python-pip|grep "install ok installed")
if [ "" == "$PKG_OK" ]; then
    echo "Installing python-pip"
    apt-get --force-yes --yes install python-pip
else
    echo "Detected python-pip installed already"
fi;

echo ""

echo "Copying SSE_Python_supervisord.conf to /etc/supervisor/conf.d"
cp SSE_Python_supervisord.conf /etc/supervisor/conf.d/SSE_Python_supervisord.conf

echo "Rereading and updating supervisorctl"
supervisorctl reread && supervisorctl update

# Install any requirements
echo "Installing project requirements via pip"
pip install -r requirements.txt
