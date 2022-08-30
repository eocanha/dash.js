#!/bin/bash

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
#NODE_VERSION=node-v8.9.4-linux-x64
NODE_VERSION=node-v14.16.1-linux-x64
export PATH="${SCRIPTDIR}/${NODE_VERSION}/bin:$PATH"

# This clones the repo in the *current* dir
git_clone() {
 cd ${SCRIPTDIR}
 REPO=https://github.com/Dash-Industry-Forum/dash.js.git
 git clone ${REPO} tmp && mv tmp/.git . && rm -rf tmp && git reset --hard
}

install_nodejs() {
 cd ${SCRIPTDIR}
 rm -rf node_modules node-v*
 TGZ=${NODE_VERSION}.tar.xz
 cd /tmp/
 curl -O https://nodejs.org/dist/v8.9.4/${TGZ}
 cd ${SCRIPTDIR}
 tar xvf /tmp/${TGZ}
 npm install
}

install_dependencies() {
 # Installation (only first time)

 # Old way:
 # npm install -g grunt-cli
 # npm install

 # New way: do nothing
 npm install -D typescript
}

build() {
 # Rebuild

 # Old way:
 #grunt debug

 # New way (for 4.1.0)
 npm run dev
}

enjoy() {
 echo
 echo "============================="
 echo " DASH.JS HAS BEEN INSTALLED! "
 echo "============================="
 echo
 echo "Assuming that ${SCRIPTDIR} is being served by your web server [somewhere], point your browser to:"
 echo "[somewhere]/samples/dash-if-reference-player/index.html?url=https://dash.akamaized.net/dash264/TestCases/1a/sony/SNE_DASH_SD_CASE1A_REVISED.mpd"
 echo
}

if [ -n "$1" ]
then
 $1
 exit 0
fi

#git_clone
#install_nodejs
install_dependencies
build
enjoy

#bash # Debugging Shell

