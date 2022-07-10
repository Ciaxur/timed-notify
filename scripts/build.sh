#!/usr/bin/env bash
ROOT_DIR=$(realpath `dirname $0`/..)
BUILD_DIR="build"
APP="timed-notify"
cd $ROOT_DIR

function get_release_name {
  BRANCH_NAME=`git rev-parse --abbrev-ref HEAD`
  BRANCH_REV=`git rev-parse HEAD`

  # Negate the matched release name, which indicats that the branch name
  # matched the pattern. ie. v1.2.3
  IS_VALID_RELEASE=$(echo $BRANCH_NAME | grep -Ev '^v(\d+\.){2}(\d+)$')
  if [[ "$IS_VALID_RELEASE" = "" ]]; then
    echo $BRANCH_NAME
  else
    echo $BRANCH_REV | grep -Eo "^(\w){6}"
  fi
}

# Clean up first.
if [[ -d $BUILD_DIR ]]; then
  echo "Cleaning up build directory."
  rm -r $BUILD_DIR
fi

# Dynamically modify the built date and version for the final binary.
OS=$(uname)
BUILD_VERSION=$(get_release_name)
BUILD_GO_FILE="internal/build/build.go"

BUILD_REGEX="s/BuildDate.*$/BuildDate=\"$(date "+%Y\/%m\/%d")\"/"
VERSION_REGEX="s/Version.*$/Version=\"$BUILD_VERSION\"/"

echo "Dynamically updating binary version."
if [[ "$OS" = "Darwin" ]]; then
  sed -i "" $BUILD_REGEX $BUILD_GO_FILE
  sed -i "" $VERSION_REGEX $BUILD_GO_FILE
  sed -i "" "s/^import.*$//" $BUILD_GO_FILE
else
  sed -i $BUILD_REGEX $BUILD_GO_FILE
  sed -i $VERSION_REGEX $BUILD_GO_FILE
  sed -i "s/^import.*$//" $BUILD_GO_FILE
fi

# Proceed with building the binary and copying over the dependent resources.
echo "Compiling binary."
go build -o ./$BUILD_DIR/$APP ./cmd/timed-notify
cp ./res/Notification.png ./$BUILD_DIR/
