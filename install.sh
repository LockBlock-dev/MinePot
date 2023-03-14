#!/bin/bash

# Prompt for sudo
sudo -v

# Create the MinePot directory
sudo mkdir /etc/minepot
# Copy the service file
sudo cp ./minepot.service /etc/systemd/system/minepot.service

# Copy the config
sudo cp ./config.json /etc/minepot/config.json

# Reload and start the service
sudo systemctl daemon-reload
sudo systemctl enable minepot.service
sudo systemctl start minepot.service
