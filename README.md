Here's an updated README that reflects the changes in your repository and clarifies the project's purpose as a CLI-based to-do application:

---

# Go Task CLI (Pre-Release)

`Go Task CLI` is a command-line tool built in Go, designed to manage tasks with statuses (`ToDo`, `In Progress`, `Done`) using JSON file storage. This project is under active development, and while the foundational code and tests are implemented, additional CLI functionality is planned for upcoming releases.

## Project Status

This pre-release version includes core structures, repository patterns, and tests. CLI functionality is partially implemented, allowing you to add, update, and delete tasks from the command line. A CI/CD pipeline ensures that every pull request (PR) to the `main` and `dev` branches, as well as commits to `main`, trigger builds and run tests.

## Features

Below is a checklist of all planned and implemented features based on project requirements:

### Core Task Management
- [x] **Task Struct and Basic Logic**: Defines a `Task` struct with fields for ID, description, status, `CreatedAt`, and `UpdatedAt`.
- [ ] **Add Task**: Allows adding new tasks from the command line, with auto-generated IDs.
- [ ] **Update Task**: Modifies task description and updates the `UpdatedAt` timestamp.
- [ ] **Delete Task**: Removes a task by its unique ID.

### Task Status Tracking
- [x] **Default Status (`ToDo`)**: Initializes tasks with a default status of `ToDo`.
- [ ] **Mark Task as In Progress**: Updates the status of a task to `In Progress`.
- [ ] **Mark Task as Done**: Updates the status of a task to `Done`.
- [ ] **List All Tasks**: Displays all tasks regardless of status.
- [x] **List Tasks by Status**: Filter tasks by status (`ToDo`, `In Progress`, `Done`).

### Data Persistence
- [ ] **JSON File Storage**: Saves tasks to a JSON file in the current directory.
- [ ] **Auto-Create JSON File**: Automatically creates the JSON file if it does not exist.
- [ ] **Handle JSON Read/Write Errors**: Gracefully manages file system and JSON errors.

### User Experience and CLI Usability
- [ ] **Positional Arguments for CLI**: Uses positional arguments to accept user inputs from the command line.
- [ ] **Error Messages for Invalid Commands**: Displays descriptive error messages for invalid actions or inputs.
- [ ] **Edge Case Handling**: Manages cases like non-existent task IDs, empty task descriptions, and invalid statuses.

### Testing and Development Workflow
- [x] **Test-Driven Development (TDD)**: Unit tests ensure reliable implementation of core features.
- [x] **CI/CD Pipeline**: Automated pipeline that builds and tests on PRs and commits to `main` and `dev` branches.

## Project Structure

The project is organized as follows:

```
go-task/
│
├── go-task.go              # Primary CLI entry point
├── task.go                 # Task struct and related logic
├── task_repository.go      # TaskRepository interface and FileTaskRepository implementation
├── file_task_repository.go # File-backed implementation of TaskRepository
├── go-task_test.go         # Unit tests for task and repository logic
├── go.mod                  # Go module file indicating dependencies
├── LICENSE                 # Project's licensing information (MIT)
└── tasks.json              # JSON file used for task data persistence
```

## Getting Started

### Installation

Clone this repository to your local environment:

```bash
git clone https://github.com/quantumburrito/go-task.git
cd go-task
go mod tidy
```

### Building the CLI

```bash
go build -o go-task
```

### Using the CLI

Once built, use `go-task` to manage tasks directly from the command line.

#### Commands

```bash
# Add a new task
./go-task add "Complete the Go project" --status "in-progress"

# List all tasks
./go-task list

# Update a task
./go-task update 1 --description "Complete CLI features" --status "done"

# Delete a task
./go-task delete 1
```

### Running Tests

This project uses test-driven development (TDD) principles, with unit tests available to validate core functionality:

```bash
go test ./...
```

## Continuous Integration

This project uses a CI/CD pipeline that:

- **Builds and Tests**: Runs automatically on every PR to `main` and `dev` branches, and on every commit to `main`.
- **Code Validation**: Ensures consistent testing to catch and prevent issues early in the development cycle.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

This README now reflects the project's current state, CLI instructions, and provides a clear roadmap of implemented and planned features. Let me know if you'd like to make further adjustments!