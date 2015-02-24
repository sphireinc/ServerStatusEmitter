# Install python2.7 if not installed
PKG_OK=$(dpkg-query -W --showformat='${Status}\n' python2.7|grep "install ok installed")
if [ "" == "$PKG_OK" ]; then
    echo "Installing python2.7"
    apt-get --force-yes --yes install python2.7
fi;

# Install python-dev if not installed
PKG_OK=$(dpkg-query -W --showformat='${Status}\n' python-dev|grep "install ok installed")
if [ "" == "$PKG_OK" ]; then
    echo "Installing python-dev"
    apt-get --force-yes --yes install python-dev
fi;

# Install python-pip if not installed
PKG_OK=$(dpkg-query -W --showformat='${Status}\n' python-pip|grep "install ok installed")
if [ "" == "$PKG_OK" ]; then
    echo "Installing python-pip"
    apt-get --force-yes --yes install python-pip
fi;

# Make the log directory for supervisord if not exists
if [ -d "/var/log/supervisord" && ! -L "/var/log/supervisord"]; then
    echo "Creating supervisord log directory"
    mkdir /var/log/supervisord
fi

# Make the log file for supervisord if not exists
if [ -f "/var/log/supervisord/supervisord.log" ]; then
    echo "Creating supervisord log file in supervisord log directory"
    touch /var/log/supervisord/supervisord.log
fi

# Install any requirements
echo "Installing project requirements via pip"
pip install -r requirements.txt