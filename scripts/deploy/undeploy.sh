#!/usr/bin/env bash

# Disable, stop, and remove unit file
sudo /bin/systemctl disable xcod.service
sudo /bin/systemctl stop xcod.service
sudo rm /etc/systemd/system/xcod.service

# Reload all unit files and reset failed
sudo /bin/systemctl daemon-reload
sudo /bin/systemctl reset-failed
