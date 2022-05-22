#!/usr/bin/env bash

sudo cp ./locr /usr/local/bin

if [ ! -d ~/.config/locr ]; then
    mkdir -pv ~/.config/locr
fi
sudo cp ./config.yaml ~/.config/locr/
