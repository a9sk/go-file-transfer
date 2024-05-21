#!/bin/bash

install_go_debian() {
    sudo apt-get update && sudo apt-get install -y golang
}

install_go_redhat() {
    sudo yum install -y golang
}

install_go_fedora() {
    sudo dnf install -y golang
}

install_go_arch() {
    sudo pacman -Syu --noconfirm go
}

install_go_macos() {
    brew install golang
}

if ! command -v go &> /dev/null; then
    echo "[!] Go is not installed on your system."
    sleep 0.3
    read -p "[?] Do you want to install Go? (y/n): " install_go

    case $install_go in
        [Yy]* )
            if [ "$(uname)" == "Linux" ]; then
                if command -v apt-get &> /dev/null; then
                    install_go_debian
                elif command -v yum &> /dev/null; then
                    install_go_redhat
                elif command -v dnf &> /dev/null; then
                    install_go_fedora
                elif command -v pacman &> /dev/null; then
                    install_go_arch    
                else
                    echo "[!] Your Linux distribution is not supported. Please install Go manually."; sleep 1; echo "[*] Exiting..."; exit 1
                fi
            elif [ "$(uname)" == "Darwin" ]; then
                install_go_macos
            else
                echo "[!] Your operating system is not supported. Please install Go manually."; sleep 1; echo "[*] Exiting..."; exit 1
            fi;;
        * )
            echo "[!] Please install Go manually and run this script again."; sleep 1; echo "[*] Exiting..."; exit 1;;
    esac
fi

go build -o cdbr cmd/main.go

if [ $? -ne 0 ]; then
    echo "[!] Build failed"
    exit 1
fi

INSTALL_DIR="/usr/local/bin"

if [ ! -d "$INSTALL_DIR" ]; then
    echo "[!] Install directory $INSTALL_DIR does not exist"
    exit 1
fi

mv cdbr "$INSTALL_DIR"

if [ $? -ne 0 ]; then
    echo "[!] Failed to move cdbr to $INSTALL_DIR"
    exit 1
fi

echo "[*] cdbr installed successfully, have fun :)"
sleep 0.5
echo '[*] Try running the cdbr command'
sleep 0.5
echo '[!] In case of any problem contact @a9sk at 920a9sk765@proton.me'