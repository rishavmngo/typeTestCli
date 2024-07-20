#!/bin/bash

# Define variables
URL="https://drive.usercontent.google.com/download?id=1xz7KFjkpt0UIy1gW2X6izodlhf49M8Fa&export=download&authuser=0&confirm=t&uuid=625eee30-d127-4432-aabf-80206bc7cb3f&at=APZUnTUIvQZfHJMEbQt3dA-eGg7x:1721459994583"
INSTALL_DIR="/usr/local/bin/Typetest-go"
TARBALL="Typetest-cli.tar.gz"
CONFIG_DIR="$HOME/.config/typeTest-go"
# Download the tarball
echo "Downloading typetest-go..."
curl -L $URL -o $TARBALL

# Create the installation directory
echo "Creating installation directory..."
sudo mkdir -p $INSTALL_DIR

# Extract the tarball to the installation directory
echo "Extracting files..."
sudo tar -xzvf $TARBALL -C $INSTALL_DIR --strip-components=1

# Make the binary executable
echo "Making the binary executable..."
sudo chmod +x $INSTALL_DIR/typeTest

# Create a symlink to the binary in /usr/local/bin
echo "Creating symlink..."
sudo ln -sf $INSTALL_DIR/typeTest /usr/local/bin/typetest

echo "Creating config directory..."
mkdir -p $CONFIG_DIR

echo "Copying configuration files"
sudo cp $INSTALL_DIR/settings.json $CONFIG_DIR/
sudo cp $INSTALL_DIR/words.json $CONFIG_DIR/

sudo chown $USER:$USER $CONFIG_DIR/settings.json
sudo chown $USER:$USER $CONFIG_DIR/words.json
# Clean up
echo "Cleaning up..."
rm $TARBALL

echo "Installation completed. You can now run 'typetest' from anywhere."
