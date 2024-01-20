#!/bin/bash

# Copy key to remote first if not done yet
# ssh-copy-id -i ~/.ssh/id_rsa.pub root@66.42.86.91

# Define local and remote paths
LOCAL_PATH="/private/var/www/websites/temp/CDN/app/src/dist/"
REMOTE_USER="root"  # Replace with your actual username on the remote server
REMOTE_IP="66.42.86.91"
REMOTE_PATH="/var/www/repos/cdn/"

echo "👷 Building...."
# Build the project
./cdn

echo "🖨️ Staring copying to $REMOTE_IP"
# SCP command to copy the folder
scp -r "$LOCAL_PATH" "$REMOTE_USER@$REMOTE_IP:$REMOTE_PATH"

echo "✅ Folder copied successfully."
