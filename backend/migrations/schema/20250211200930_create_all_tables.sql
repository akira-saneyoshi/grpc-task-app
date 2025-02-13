-- データベースを作成するSQL (存在しない場合のみ作成)
CREATE DATABASE IF NOT EXISTS mydatabase;

-- 作成したデータベースを使用するSQL
USE mydatabase;

-- usersテーブルを作成するSQL
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY unique_email (email)
);

-- tasksテーブルを作成するSQL
CREATE TABLE IF NOT EXISTS tasks (
    id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('pending', 'in_progress', 'completed') DEFAULT 'pending',
    due_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE KEY unique_user_task_title (user_id, title)
);

-- usersテーブルのサンプルデータを挿入するSQL
INSERT INTO users (id, name, email, password, is_active) VALUES
('user1', '田中太郎', 'testuser1@example.com', '$2a$10$YfHxWNfL8Ba2ltl6TRHMVuN0WPXxAuB5L7w1Y0jqaFcn2bUDoUq9W', TRUE),
('user2', '田中次郎', 'testuser2@example.com', '$2a$10$YfHxWNfL8Ba2ltl6TRHMVuN0WPXxAuB5L7w1Y0jqaFcn2bUDoUq9W', TRUE),
('user3', '田中無効', 'inactiveuser@example.com', '$2a$10$YfHxWNfL8Ba2ltl6TRHMVuN0WPXxAuB5L7w1Y0jqaFcn2bUDoUq9W', FALSE);

-- tasksテーブルのサンプルデータを挿入するSQL
INSERT INTO tasks (id, user_id, title, description, status, due_date) VALUES
('task1', 'user1', 'タスク 1', 'Description for Task 1', 'pending', '2025-03-10'),
('task2', 'user1', 'タスク 2', 'Description for Task 2', 'in_progress', '2025-03-15'),
('task3', 'user2', 'タスク 3', 'Description for Task 3', 'completed', '2025-03-05'),
('task4', 'user1', 'タスク 4', 'Description for Task4', 'pending', '2025-04-01');
