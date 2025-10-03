import os
import subprocess
    
def compile_binaries():
    '''Compiles the Go files into binaries and returns the compiled binaries' paths.'''
    compiled_binaries = []
    try:
        main_path = './'
        for subdir in os.listdir(main_path):
            subdir_path = os.path.join(main_path, subdir)
            if not os.path.isdir(subdir_path):
                continue
            for file in os.listdir(subdir_path):
                if file.endswith('.go'):
                    subprocess.run(['go', 'build', file], cwd=subdir_path)
                    executable_path = file[:-3]  # Remove .go extension
                    if os.path.exists(os.path.join(subdir_path, executable_path)):
                        compiled_binaries.append(os.path.join(subdir_path, executable_path))

        return compiled_binaries
    except Exception as e:
        print(f"Error compiling binaries: {e}")
        return []
    
def add_alias(binary):
    '''Adds an alias for the given binary and writes it to the .zshrc file, returning the alias name.'''
    pass

def main():
    print('Welcome to Promptly Setup!')
    print('Compiling Go binaries...')
    compiled_binaries = compile_binaries()

    for binary in compiled_binaries:
        alias = add_alias(binary)
        print(f'Alias for binary: {binary} -> {alias}')
    
    print('Setup complete!')

if __name__ == '__main__':
    main()
