#!/bin/bash

# Define variables
URL="https://drive.usercontent.google.com/download?id=1Y3vGxnnRHedFEpQFeYz9Ld8P7vC5HpWq&export=download&authuser=0&confirm=t&uuid=da44af97-6716-45c2-a442-a7d7d0f85fa6&at=APZUnTVwOCpPVT2GUiQJmymbbaLz:1721448445018"
INSTALL_DIR="/usr/local/bin/Typetest-go"
TARBALL="Typetest-cli.tar.gz"

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

# Clean up
echo "Cleaning up..."
rm $TARBALL

echo "Installation completed. You can now run 'typetest' from anywhere."
