#!/usr/bin/env bash

function log {
    echo -e "$1[.]${RESET} $2"
}

function success_log {
    echo -e "$1[^]${RESET} $2"
}

function input {
    echo -e -n "$1[?]${RESET} $3: "
    read -r "$2"
}

function chap_log {
    echo
    echo -e "${BOLD}$1 ${WHITE}$2${RESET}"
}

GREEN="\x1B[32m"
WHITE="\x1B[37m"
RED="\x1B[31m"
CYAN="\x1B[36m"
BOLD="\x1B[1m"
MAGENTA="\x1B[35m"
BLUE="\x1B[34m"
RESET="\x1B[0m"

log "$CYAN" "Welcome to the installation script for ${BOLD}IndiePKG${RESET}!"

chap_log "$MAGENTA=>" "Checking for dependencies"

if git --version >/dev/null 2>&1; then
    log "$CYAN" "Git installed."
else
    log "$RED" "Git not installed. Please install git before continuing."
    exit 1
fi

if go version >/dev/null 2>&1; then
    log "$CYAN" "Go installed."
else
    log "$RED" "Go not installed. Please install Go before continuing."
    exit 1
fi

if make --version >/dev/null 2>&1; then
    log "$CYAN" "Make installed."
else
    log "$RED" "Make not installed. Please install Make before continuing."
    exit 1
fi

chap_log "$MAGENTA=>" "Installing IndiePKG"

chap_log "$BLUE==>" "Cloning source code"

log "$CYAN" "Cloning source code..."
mkdir "$HOME/.indiepkg/"
git clone https://github.com/talwat/indiepkg.git "$HOME/.indiepkg/src"

chap_log "$BLUE==>" "Compiling source code"

log "$CYAN" "Compiling source code..."
cd "$HOME/.indiepkg/src" || exit 1
make

chap_log "$BLUE==>" "Installing"

log "$CYAN" "Installing..."
make install


chap_log "$BLUE==>" "Running indiepkg setup"
"$HOME/.local/bin/indiepkg" setup

chap_log "$GREEN=>" "Success"
success_log "$GREEN" "Installation complete!"