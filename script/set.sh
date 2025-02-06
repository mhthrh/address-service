#!/usr/bin/bash
env_path="A3PATH"
env_value="/Users/mohsen/Projects/Golang/x-bank"

shell_config_file="$(echo $SHELL | grep -o 'bash\|zsh')"

if [ "$shell_config_file" = "bash" ]; then
    shell_config_file="$HOME/.bashrc"
elif [ "$shell_config_file" = "zsh" ]; then
    shell_config_file="$HOME/.zshrc"
else
    echo "Unsupported shell: $shell_config_file"
    exit 1
fi

if ! grep -q "export $env_path=" "$shell_config_file"; then
  echo "export $env_path=\"$env_value\"" >> "$shell_config_file"
else
  echo "Variable already exists in $shell_config_file"
fi

source "$shell_config_file"

