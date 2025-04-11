# Notion Task CLI

A command-line tool for quickly adding tasks to a Notion database via an interactive chat interface. Built with Go.

## Overview

Notion Task CLI provides a simple, conversational way to add tasks to your designated Notion database directly from your terminal. No need to switch contexts – just type your task details in a natural way.

<!-- Optional: Add a GIF/Screenshot demonstrating the chat interface here -->
<!-- ![Demo GIF](link/to/your/demo.gif) -->

## Features

*   **Interactive Chat Mode:** Add tasks conversationally.
*   **Simple Syntax:** Use keywords like `due`, `priority`, and `status`.
*   **Natural Language Date Parsing:** Understands terms like `today`, `tomorrow`, `next monday`, `YYYY-MM-DD`, etc.
*   **Configurable:** Connects to your specific Notion database via an integration token.

## Installation

1.  **Prerequisites:** Ensure you have [Go](https://go.dev/doc/install) (version 1.x or later) installed.
2.  **Clone:** Clone this repository:
    ```bash
    git clone https://github.com/justKevv/notion-task-cli.git 
    cd notion-task-cli
    ```
3.  **Build:** Build the executable:
    ```bash
    go build -o notion-task
    ```
    *(This creates an executable named `notion-task` in the current directory)*

## Configuration

Before running the tool, you need to configure your Notion API credentials:

1.  **Create a `.env` file** in the root directory of the project (`notion-task-cli/`).
2.  **Add the following variables** to the `.env` file:
    ```dotenv
    NOTION_TOKEN=your_notion_integration_secret_here
    NOTION_DATABASE_ID=your_notion_database_id_here
    ```
3.  **Obtain the values:**
    *   **`NOTION_TOKEN`**:
        *   Go to [Notion's My Integrations page](https://www.notion.so/my-integrations).
        *   Click "Create new integration". Give it a name (e.g., "CLI Task Adder").
        *   Ensure it's associated with the correct workspace.
        *   Copy the "Internal Integration Secret" – this is your token.
    *   **`NOTION_DATABASE_ID`**:
        *   Navigate to the Notion database you want to add tasks to.
        *   Share the database with the integration you just created (Click "..." -> "Add connections" -> Find your integration).
        *   Copy the Database ID from the URL. It's the long string of characters between your workspace name and the `?v=` part.
          *Example URL:* `https://www.notion.so/yourworkspace/`**`abcdef1234567890abcdef1234567890`**`?v=...` -> The bold part is the ID.

## Usage

1.  **Run the chat command** from the project directory:
    ```bash
    ./notion-task chat
    ```
2.  **Enter the interactive chat mode.** You'll see a `>` prompt.
3.  **Add tasks** using the format:
    ```plaintext
    add [task name] due [date] priority [level] status [status]
    ```
    *   Keywords (`due`, `priority`, `status`) are optional.
    *   The order of keywords generally doesn't matter after the task name.
4.  **Examples:**
    ```plaintext
    > add Finish project report due tomorrow priority high status In Progress
    > add Call Mom due next friday
    > add Submit expense report status Done priority medium
    > add Buy milk due 2025-05-10
    ```
5.  **Exit:** Type `exit` or `quit` to leave the chat mode.

## Developer Note: Learning Go

This project was developed as a learning exercise for the Go programming language. While functional, the primary goal was educational. Code structure and practices might reflect this learning process. Feedback and suggestions are welcome via GitHub Issues!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
