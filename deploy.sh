#!/bin/bash

# Copy key to remote first if not done yet
# ssh-copy-id -i ~/.ssh/id_rsa.pub root@66.42.86.91

# Define local and remote paths
MIN_CSS="/Users/yodarango/Desktop/repos/cdn/app/src/dist/"
ICONS="/Users/yodarango/Desktop/repos/cdn/app/src/icons/"
REMOTE_USER="root"  # Replace with your actual username on the remote server
REMOTE_IP="66.42.86.91"
REMOTE_PATH="/var/www/repos/cdn/"

if [ "$#" -ne 1 ]; then
    echo "Please provide a commit message"
    exit 1
fi

# The commit message is the first argument to the script
COMMIT_MESSAGE="$1"

echo "📝 Committing and pushing changes..."

git add .
git commit -m "$COMMIT_MESSAGE"
git push 

echo "👷 Building...."
# Build the project
./cdn

echo "🖨️ Staring copying to $REMOTE_IP"
# SCP command to copy the folder
scp -r "$MIN_CSS" "$REMOTE_USER@$REMOTE_IP:$REMOTE_PATH"
rsync -avz "$ICONS" "$REMOTE_USER@$REMOTE_IP:${REMOTE_PATH}/dist/icons/"

echo "✅ Deploy executed successfully."
