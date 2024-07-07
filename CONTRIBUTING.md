# Contributing to terminal-kanban

First off, thank you for considering contributing to our project! Here are some guidelines to help you get started.

## Code of Conduct

This project and everyone participating in it is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How to Contribute

### Fork and Clone

1. Fork the repository by clicking the "Fork" button.
2. Clone the repository to your local machine:
   ```sh
   git clone https://github.com/AbilityJLR/terminal-kanban.git
   cd terminal-kanban
   ```
3. Set localhost for the development environment in the `etc/hosts` file:

   For macOS and Linux:
   1. Use `sudo vi /etc/hosts`.
   2. Add the following line:
   ```
   127.0.0.1 kanbanterm.dev
   ```
   3. Exit by typing `:wq` and pressing Enter.

   For Windows:
   1. Open your text editor in Administrator mode.
   2. In the text editor, open `C:\Windows\System32\drivers\etc\hosts`.
   3. Add the following line:
      ```
      127.0.0.1 kanbanterm.dev
      ```
   4. Save the changes.
      
### Pull Request Process
1. Create a new branch for your changes:
   ```
   git checkout -b feature/your-feature
   ```
2. Run all processes automatically:
   ### ⚠️ Simply run
   ```
   skaffold dev
   ```
   in your terminal.

3. Make your changes and commit them with a meaningful message:
   ```
   git commit -m "Add feature X"
   ```
4. Push your changes to your fork:
   ```
   git push origin feature/your-feature
   ```
5. Open a pull request on GitHub.

## Reporting Issues
If you encounter any issues, please report them here with detailed information.

## Commit Message Guidelines
Use the present tense ("Add feature" not "Added feature").
Reference issues and pull requests liberally after the first line.

## License
By contributing, you agree that your contributions will be licensed under the project's MIT License.
