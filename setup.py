import os
import subprocess
    

print('Welcome to Promptly Setup!')
print('Compiling binary...')

subprocess.run(['go', 'build', '-o', 'promptly', 'main.go'], check=True)

with open(os.path.expanduser('~/.zshrc'), 'a') as zshrc:
    zshrc.write('\nalias pmt="{}"\n'.format(os.path.abspath('promptly')))

print('Setup complete!')

