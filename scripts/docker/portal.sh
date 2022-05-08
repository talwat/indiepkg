#!/usr/bin/env bash
# DO NOT RUN THIS SCRIPT ON YOUR HOST MACHINE, THIS IS INTENDED TO BE RAN BY A DOCKER CONTAINER

if [ "$1" = "run" ]; then
    "${@:2}"
else
    "/indiepkg" "-y" "${@:1}"
fi