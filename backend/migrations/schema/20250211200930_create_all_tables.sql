-- データベースを作成するSQL (存在しない場合のみ作成)
CREATE DATABASE IF NOT EXISTS mydatabase;

-- 作成したデータベースを使用するSQL
USE mydatabase;

-- usersテーブルを作成するSQL
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
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
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('pending', 'in_progress', 'completed') DEFAULT 'pending',
    due_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE KEY unique_user_task_title (user_id, title)
);

-- usersテーブルのサンプルデータを挿入するSQL (オプション)
INSERT INTO users (name, email, password, is_active) VALUES
('John Doe', 'john.doe@example.com', 'password123', TRUE),   -- is_active を TRUE に設定 (有効)
('Jane Smith', 'jane.smith@example.com', 'password456', FALSE);  -- is_active を FALSE に設定 (無効)

-- tasksテーブルのサンプルデータを挿入するSQL (オプション)
INSERT INTO tasks (user_id, title, description, status, due_date) VALUES
(1, 'プロジェクトAのタスク1', '詳細な説明1', 'in_progress', '2024-12-31'),
(1, 'プロジェクトAのタスク2', '詳細な説明2', 'pending', '2025-01-15'),
(2, 'プロジェクトBのタスク1', '詳細な説明3', 'completed', '2024-11-30');
