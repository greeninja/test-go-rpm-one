#!/bin/bash

major=$1
minor=$2

echo "Building greenweb go rpm - v$major-$minor"
podman build --build-arg MAJOR=$major --build-arg MINOR=$minor -t greeninja/greenweb:build . -f Dockerfile.build

echo "Running build container"
podman container create --name extract greeninja/greenweb:build
echo "Copying Greenweb rpm"
podman container cp extract:/build/greenweb-$major-$minor.x86_64.rpm ./rpms/
echo "Removing build container"
podman container rm -f extract
