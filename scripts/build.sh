#!/usr/bin/env bash
ROOT_DIR=$(realpath `dirname $0`/..)
BUILD_DIR="build"
APP="timed-notify"
cd $ROOT_DIR

function get_release_name {
  BRANCH_REV=`git rev-parse HEAD`

  # Get a list of all known tags.
  MATCHING_TAG=`git show-ref --tags | grep  "$BRANCH_REV"`
  RELEASE_TAG=`echo $MATCHING_TAG | grep -oE 'v([0-9]+\.){2}([0-9]+)$'`

  # Check on matching pattern, which indicats that the branch name
  # matched the pattern. ie. v1.2.3
  if [[ "$RELEASE_TAG" != "" ]]; then
    echo $RELEASE_TAG
  else
    echo dev-$(echo $BRANCH_REV | grep -Eo "^(\w){8}")
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
if [ $? != 0 ]; then echo "Failed to build go package!" && exit 1; fi
cp ./res/Notification.png ./$BUILD_DIR/
