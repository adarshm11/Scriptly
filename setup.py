import os
import subprocess
    

print('Welcome to Promptly Setup!')
print('Compiling binary...')

try:
    subprocess.run(['go', 'build', '-o', 'promptly', 'main.go'], check=True)
    subprocess.run(['chmod', '+x', 'promptly'], check=True)
    print('Binary compiled successfully!')

    with open(os.path.expanduser('~/.zshrc'), 'a') as zshrc:
        zshrc.write(f'\nalias pmt="{os.path.abspath("promptly")}"\n')

    print('Setup complete! Reload your terminal or run "source ~/.zshrc" to apply changes. Then you can call Promptly commands using the alias "pmt".')

except subprocess.CalledProcessError as e:
    print('An error occurred while compiling the binary:', e)

except Exception as e:
    print('An unexpected error occurred:', e)