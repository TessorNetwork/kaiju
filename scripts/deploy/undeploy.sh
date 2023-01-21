#!/usr/bin/env bash

# Disable, stop, and remove unit file
sudo /bin/systemctl disable kaijud.service
sudo /bin/systemctl stop kaijud.service
sudo rm /etc/systemd/system/kaijud.service

# Reload all unit files and reset failed
sudo /bin/systemctl daemon-reload
sudo /bin/systemctl reset-failed
