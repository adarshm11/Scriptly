# Scriptly

Navigating the Mac command line is tough, even for some intermediate and advanced developers. There's a vast catalog of commands to traverse through gigantic file systems, and it's difficult to remember the functionality of most, let alone all, of these commands. Scriptly aims to fix this problem by providing you with a super easy, customizable, sleek command line interface that makes your Mac's scripting experience a breeze. With concise aliases, robust functionality, and a transparent setup, Scriptly enhances your command line experience so that you never have to worry about forgetting a shell command again.

## Functionalities

To call a Scriptly function, simply use the prefix `sc`. Scriptly commands include, but are not limited to:
1. **System Checks**. Scriptly allows you to check what items are taking up storage on your Mac that may be difficult to find.
2. **Docker Control**. The Docker command line can be verbose, but Scriptly easily abstracts the process of viewing containers and images, disk usage, and cleaning up your Docker items.
3. **Git**. With so many version control commands, Scriptly creates sub-aliases for those obscure Git commands that are so useful, but so frustrating to memorize.

For a full list of Scriptly capabilities, simply run `sc help`. 

## Setup
1. Ensure you have Go installed on your system. If not: visit `https://go.dev` to install.
2. Clone this repository: `git clone https://github.com/adarshm11/Scriptly.git`
3. From the main directory, run `python setup.py` to set up the CLI.

And you're all set! The setup script will compile all of the Go binaries, and set aliases for these inside your .zshrc file. You can customize these aliases to whatever you like by editing the `setup.py` file.