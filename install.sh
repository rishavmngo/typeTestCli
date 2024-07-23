#!/bin/bash

trap 'echo "Script interrupted. Exiting..."; exit 1;' INT
INSTALL_DIR="/usr/local/bin/Typetest-go"

TARBALL="Typetest-cli.tar.gz"
CONFIG_DIR="$HOME/.config/typeTest-go"
GREEN='\033[0;32m'
RED='\033[0;31m'
RESET='\033[0m'

DEFAULT_ARG1=""
ARG1=${1:-$DEFAULT_ARG1}
./uninstall.sh

print_green() {
  echo -e "${GREEN}$1${RESET}"
}

print_red() {
  echo -e "${RED}$1${RESET}"
}

if [ "$ARG1" == "build" ]; then

  if [ -f "$TARBALL" ]; then
    rm $TARBALL
  fi
  ./build.sh
fi

if [ -f "$TARBALL" ]; then
  print_green "Found a tarball in current directory"
else
  print_red "Can't able to find tarball in current directory"
  exit 1
fi

# Create the installation directory
print_green "Creating installation directory..."
sudo mkdir -p $INSTALL_DIR

# Extract the tarball to the installation directory
print_green "Extracting files..."
sudo tar -xzvf $TARBALL -C $INSTALL_DIR --strip-components=1

# Make the binary executable
print_green "Making the binary executable..."
sudo chmod +x $INSTALL_DIR/typeTest

# Create a symlink to the binary in /usr/local/bin
print_green "Creating symlink..."
sudo ln -sf $INSTALL_DIR/typeTest /usr/local/bin/typetest

print_green "Creating config directory..."
mkdir -p $CONFIG_DIR

print_green "Copying configuration files"
sudo cp $INSTALL_DIR/settings.json $CONFIG_DIR/
sudo cp $INSTALL_DIR/words.json $CONFIG_DIR/

sudo chown $USER:$USER $CONFIG_DIR/settings.json
sudo chown $USER:$USER $CONFIG_DIR/words.json
# Clean up
print_green "Cleaning up..."
# rm $TARBALL
rm typeTest

print_green "Installation completed. You can now run 'typetest' from anywhere."
