#!/bin/bash

TARBALL="Typetest-cli.tar.gz"
INSTALL_DIR="/usr/local/bin/Typetest-go"
CONFIG_DIR="$HOME/.config/typeTest-go"
INSTALL_SYMBOLIC="/usr/local/bin/typetest"

GREEN='\033[0;32m'
RESET='\033[0m'

print_green() {

  echo -e "${GREEN}$1${RESET}"
}
IS_INSTALLED=0

if [ -f "$INSTALL_SYMBOLIC" ]; then
  sudo rm "$INSTALL_SYMBOLIC"
  print_green "Deleting the binary"
  IS_INSTALLED=1
fi

if [ -d "$INSTALL_DIR" ]; then
  sudo rm -rf "$INSTALL_DIR"
  print_green "Deleting the Installation directory"
fi

if [ -d "$CONFIG_DIR" ]; then
  sudo rm -rf "$CONFIG_DIR"
  print_green "Deleting the Configuration directory"
fi

if [ $IS_INSTALLED -eq 1 ]; then
  print_green "Typetest completely removed"
else
  print_green "Typetest not installed"
fi
