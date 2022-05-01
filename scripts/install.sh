#!/usr/bin/env bash
#shellcheck disable=SC2059


log() {
    echo -e "$1[.]${RESET} $2"
}

success_log() {
    echo -e "$1[^]${RESET} $2"
}

input() {
    echo -e -n "$1[?]${RESET} $3: "
    read -r "$2"
}

chap_log() {
    echo
    echo -e "${BOLD}$1 ${WHITE}$2${RESET}"
}

GREEN="\033[32m"
WHITE="\033[37m"
RED="\033[31m"
CYAN="\033[36m"
BOLD="\033[1m"
MAGENTA="\033[35m"
BLUE="\033[34m"
RESET="\033[0m"

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