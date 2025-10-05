import os
import subprocess
    

print('Welcome to Promptly Setup!')
print('Compiling binary...')

try:
    subprocess.run(['go', 'build', '-o', 'promptly', 'main.go'], check=True)

    with open(os.path.expanduser('~/.zshrc'), 'a') as zshrc:
        zshrc.write(f'\nalias pmt="{os.path.abspath("promptly")}"\n')

    print('Setup complete! You can now call Promptly commands using the alias "pmt".')

except subprocess.CalledProcessError as e:
    print('An error occurred while compiling the binary:', e)

except Exception as e:
    print('An unexpected error occurred:', e)