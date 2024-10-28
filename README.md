# Go Task CLI (Pre-Release)

`Go Task CLI` is a command-line interface tool designed in Go to manage tasks with statuses (ToDo, In Progress, Done) using a JSON storage approach. This app is under active development, following a Test-Driven Development (TDD) approach. Currently, only unit tests are available to verify core functionalities, with a full CLI implementation planned for future releases.

## Current Status

This is a pre-release version, intended primarily for testing purposes. The CLI features are not yet operational, but the foundational code and tests for handling tasks are in place.

## Features

In the current state, `Go Task CLI` includes:

- **Task Struct and Basic Logic**: Core logic for task creation, including ID generation, timestamps (`createdAt` and `updatedAt`), and status (`ToDo` by default).
- **JSON Data Structure**: Planned JSON storage design to save and load task data in future versions.
- **Test-Driven Development**: Unit tests are implemented to ensure key functionalities align with requirements.
  
## Requirements

The application will eventually fulfill the following requirements:

- **Task Actions**: Allow users to add, update, delete, and mark tasks as "in progress" or "done".
- **Task Status Tracking**: List tasks filtered by status (`ToDo`, `In Progress`, `Done`).
- **Data Persistence**: Tasks will be saved to a JSON file, which will be created automatically if not present.
- **Error Handling**: Graceful handling of invalid inputs and missing files.
  
## Project Structure

```
go-task/
│
├── go-task.go           # Core logic with Task struct (pre-release state)
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

The tests cover:

- **Task Creation**: Ensures each task has unique IDs, default status as `ToDo`, and proper timestamps.
- **Structured Task List Setup**: Initializes a list of tasks to verify core list handling.

## TODO: Planned Features

These features will be developed in upcoming releases, based on the initial requirements:

1. **CLI Task Management**:
   - Add CLI commands to add, update, delete, and list tasks based on status.

2. **JSON Storage Integration**:
   - Implement JSON read/write functionality for data persistence across sessions.

3. **Error and Edge Case Handling**:
   - Manage non-existent task IDs, missing JSON files, and invalid inputs with user-friendly error messages.

4. **Enhanced Status Commands**:
   - Extend commands to set tasks as "in progress" or "done" and filter lists based on these statuses.

5. **Full Task Lifecycle Testing**:
   - Expand tests to cover edge cases and the complete task management lifecycle.

## License

This project is licensed under the MIT License.

---

This README sets clear expectations for current functionality and the planned roadmap. Let me know if this matches your needs or if there’s anything more to add!
