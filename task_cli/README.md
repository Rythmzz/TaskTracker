# Task CLI - Manage tasks from command line running on windows
# Project URL: https://roadmap.sh/projects/task-tracker

# Step 1: Build program

- go build -o task-cli.exe .

# Step 2: Add current directory to PATH (Run once)

# According to windows users

- $env:PATH += ";$pwd"

# According to linux users

- echo 'export PATH="$PATH:'$(pwd)'"' >> ~/.bashrc
- source ~/.bashrc

# Step 3: Run CMD 

- task-cli add "Go buy food"