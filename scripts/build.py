#!/usr/bin/env python3

import subprocess
import os

cmd_out = subprocess.getoutput("go tool dist list")
split = cmd_out.split("\n")
supported_os = ["linux", "darwin"]
supported_arch = ["amd64", "386", "arm", "arm64", "ppc64", "s390x"]

for i in split:
    i_split = i.split("/")
    if i_split[0] in supported_os and i_split[1] in supported_arch:
        os.system(
            f"env GOOS={i_split[0]} GOARCH={i_split[1]} go build -o output/indiepkg-{i_split[0]}-{i_split[1]}"
        )
