# Promptly

Navigating the command line is tough, even for some intermediate and advanced developers. There's a vast catalog of commands to traverse through gigantic file systems, and it's difficult to remember the functionality of most, let alone all, of these commands. Promptly aims to fix this problem by providing you with a super easy, customizable, sleek command line interface that makes your scripting experience a breeze. With concise aliases, robust functionality, and a transparent setup, Promptly enhances your command line experience so that you never have to worry about forgetting a shell command again.

## Setup
1. Ensure you have Go installed on your system. If not: visit `https://go.dev` to install.
2. Clone this repository: `git clone https://github.com/adarshm11/Promptly.git`
3. From the main directory, run `python setup.py` to set up the CLI.

And you're all set! The setup script will compile all of the Go binaries, and set aliases for these inside your .bashrc or .zshrc file. You can customize these aliases to whatever you like by editing the `setup.py` file.