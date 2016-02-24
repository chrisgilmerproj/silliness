#!/bin/bash

wget https://s3.amazonaws.com/drone.deploy.map.engine/example.zip
mkdir -p images/
unzip example.zip
mv *.jpg images/
rm -rf __MACOSX
rm example.zip
