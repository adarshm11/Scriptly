import os
import subprocess
    

print('Welcome to Scriptly Setup!')
print('Compiling binary...')

try:
    subprocess.run(['go', 'build', '-o', 'scriptly', 'main.go'], check=True)
    subprocess.run(['chmod', '+x', 'scriptly'], check=True)
    print('Binary compiled successfully!')

    with open(os.path.expanduser('~/.zshrc'), 'a') as zshrc:
        zshrc.write(f'\nalias sc="{os.path.abspath("scriptly")}"\n')

    print('Setup complete! Reload your terminal or run "source ~/.zshrc" to apply changes. Then you can call Scriptly commands using the alias "sc".')

except subprocess.CalledProcessError as e:
    print('An error occurred while compiling the binary:', e)

except Exception as e:
    print('An unexpected error occurred:', e)