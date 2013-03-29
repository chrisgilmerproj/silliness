#! /bin/bash

export ARCHFLAGS="-arch i386 -arch x86_64"
export CC=clang

brew install rabbitmq

virtualenv env
ln -s env/bin/activate
source activate

pip install -r requirements.txt
