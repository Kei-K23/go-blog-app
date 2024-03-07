# Go Blog App

This is a simple blog application built with Go (Golang) and Echo framework. It allows users to create, view, edit, and delete blog posts.

## Features

- Create blog posts with a title, description, and content.
- View individual blog posts.
- Edit existing blog posts.
- Delete blog posts.
- Upload and serve images for blog posts.

## Requirements

- Go (Golang)
- Echo framework
- SQL database (SQLite, MySQL, PostgreSQL, etc.)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Kei-K23/go-blog-app.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-blog-app
   ```

3. Create a `.env` file in the project root and specify your database connection URL:

   ```plaintext
   DB_URL=<your_database_connection_url>
   ```

4. Install dependencies:

   ```bash
   go mod tidy
   ```

5. Build and run the application:

   ```bash
   go run main.go
   ```

## Database Setup

The application uses a SQL database to store blog posts. Make sure you have a SQL database server installed and running.

By default, the application uses Turso (SQLite). If you want to use a different database (MySQL, PostgreSQL, etc.), you'll need to modify the database driver and connection URL in the code.

## Usage

- Access the application by visiting `http://localhost:8080` in your web browser.
- Click on the "Create New Blog" button to create a new blog post.
- Click on a blog post title to view its details.
- Click on the "Edit" button to edit a blog post.
- Click on the "Delete" button to delete a blog post.

Feel free to contribute to the project and improve its features!
