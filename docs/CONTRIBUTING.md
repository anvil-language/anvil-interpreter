# contributing

>[!NOTE]
> You will need to install go for either [windows](https://go.dev/dl/go1.27.1.windows-amd64.msi) or linux `sudo apt install golang-go`

## Forking, Committing, and Pushing

### Fork the Repository
1. Click **Fork** on the repository page.  
2. Clone your fork locally:

```sh
git clone https://github.com/<your-username>/anvil-interpreter.git
cd anvil-interpreter
```

---

### Create a Branch
Always create a new branch for your work:

```sh
git checkout -b feature/<short-description>
```

Examples:  
`feature/new-parser`, `fix/typo-in-readme`, `improve/error-handling`

---

### Make Your Changes
Implement your changes in small, focused commits.

Then build the source code with the following command and go installed:
```sh
go build -o bin/anvil ./cmd/anvil
```
Then try out your features in the repl with 
```sh
./bin/anvil
```
or make a `.an` file and run `./bin/anvil example.an`

---

### Commit Your Work
Stage and commit your changes with a clear message:

```sh
git add .
git commit -m "Describe the change made"
```

Commit messages should be concise and in present tense.

---

### Push Your Branch
Push your branch to your fork:

```sh
git push origin feature/<short-description>
```

---

### Open a Pull Request
Go to your fork on GitHub and open a **Pull Request** targeting the main repository.  
Include a short summary of your changes and any relevant context.
