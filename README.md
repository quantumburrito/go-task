Here’s a comprehensive README draft for your **Go Task CLI** project, incorporating the requirements doc, a feature checklist, a release schedule, and notes on the CI pipeline. This format provides clarity on current and future functionality, making it ideal for contributors and users to track development progress.

---

# Go Task CLI (Pre-Release)

`Go Task CLI` is a command-line interface tool built in Go, designed to help manage tasks with statuses (ToDo, In Progress, Done) using a JSON storage file for persistence. This project is under active development, and the current version focuses on foundational code and testing.

## Project Status

This pre-release version includes only core structures and tests. CLI functionality is planned for upcoming releases. A CI pipeline is in place to ensure that every pull request (PR) to `main` and `dev` branches, as well as commits to `main`, trigger builds and run tests.

## Features

The following is a checklist of all planned and implemented features based on the requirements:

### Core Task Management
- [x] **Task Struct and Basic Logic**: Create a `Task` struct with fields for ID, description, status, created_at, and updated_at.
- [ ] **Add Task**: Add new tasks from the command line with an auto-generated ID.
- [ ] **Update Task**: Modify task description and update the `updatedAt` timestamp.
- [ ] **Delete Task**: Remove a task by its unique ID.

### Task Status Tracking
- [x] **Default Status (`ToDo`)**: Initialize tasks with a default status of `ToDo`.
- [ ] **Mark Task as In Progress**: Update the status of a task to "In Progress".
- [ ] **Mark Task as Done**: Update the status of a task to "Done".
- [ ] **List All Tasks**: Display all tasks regardless of status.
- [ ] **List Tasks by Status**: Filter tasks by status (`ToDo`, `In Progress`, `Done`).

### Data Persistence
- [ ] **JSON File Storage**: Save tasks to a JSON file in the current directory.
- [ ] **Auto-Create JSON File**: Automatically create the JSON file if it does not exist.
- [ ] **Handle JSON Read/Write Errors**: Gracefully handle file system and JSON errors.

### User Experience and CLI Usability
- [ ] **Positional Arguments for CLI**: Use positional arguments to accept user inputs from the command line.
- [ ] **Error Messages for Invalid Commands**: Display descriptive error messages for invalid actions or inputs.
- [ ] **Edge Case Handling**: Manage cases like non-existent task IDs, empty task descriptions, and invalid statuses.

### Testing and Development Workflow
- [x] **Test-Driven Development (TDD)**: Write unit tests for key functions to ensure reliable implementation.
- [x] **CI/CD Pipeline**: Automated pipeline that builds and tests on PRs and commits to `main` and `dev` branches.

## Project Structure

The project is organized as follows:

```
go-task/
│
├── go-task.go           # Primary codebase with Task struct and task logic
├── go-task_test.go      # Unit tests to verify foundational logic
├── go.mod               # Go module file indicating dependencies
├── LICENSE              # Project's licensing information (MIT)
└── tasks.json           # Planned JSON storage file for task data (not yet functional)
```

## Getting Started

### Installation

Clone this repository to your local environment:

```bash
git clone https://github.com/quantumburrito/go-task.git
cd go-task
```

Currently, only testing functionality is available, as the CLI is not yet operational.

### Running Tests

Since this pre-release focuses on TDD, you can run the included unit tests to validate the basic structure and logic:

```bash
go test ./...
```


## Continuous Integration

This project uses a CI/CD pipeline that:

- **Builds and Tests**: Runs automatically on every pull request (PR) to the `main` and `dev` branches, and on every commit to `main`.
- **Code Validation**: Ensures consistent testing to detect and prevent issues early in the development cycle.

## License

This project is licensed under the MIT License.

---

This README provides a clear overview of current progress and future milestones, with a checklist to easily track feature completion. Let me know if there are any adjustments you'd like to make!