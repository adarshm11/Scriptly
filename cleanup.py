import os
import subprocess

zshrc_path = os.path.expanduser("~/.zshrc")

with open(zshrc_path, "r") as f:
    lines = f.readlines()

with open(zshrc_path, "w") as f:
    for line in lines:
        if not line.strip().startswith("alias pmt="):
            f.write(line)

print('Cleanup complete! Reload your terminal or run "source ~/.zshrc" to apply changes.')