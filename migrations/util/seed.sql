INSERT INTO users (username, email, created_at, updated_at)
VALUES
    ('john_doe', 'john.doe@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('jane_smith', 'jane.smith@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('bob_jones', 'bob.jones@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO posts (title, content, user_id, created_at, updated_at)
VALUES
    ('First Post', 'This is the content of the first post.', (SELECT id FROM users WHERE username = 'john_doe'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Second Post', 'Content for the second post.', (SELECT id FROM users WHERE username = 'jane_smith'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Third Post', 'Content for the third post.', (SELECT id FROM users WHERE username = 'bob_jones'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO comments (content, user_id, post_id, created_at, updated_at)
VALUES
    ('Great post!', (SELECT id FROM users WHERE username = 'jane_smith'), (SELECT id FROM posts WHERE title = 'First Post'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('I agree!', (SELECT id FROM users WHERE username = 'bob_jones'), (SELECT id FROM posts WHERE title = 'First Post'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Nice content.', (SELECT id FROM users WHERE username = 'john_doe'), (SELECT id FROM posts WHERE title = 'Second Post'), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO ratings (value, user_id, post_id, comment_id, created_at, updated_at)
VALUES
    (ROUND(RANDOM() * 4) + 1, (SELECT id FROM users ORDER BY RANDOM() LIMIT 1), (SELECT id FROM posts ORDER BY RANDOM() LIMIT 1), NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (ROUND(RANDOM() * 4) + 1, (SELECT id FROM users ORDER BY RANDOM() LIMIT 1), (SELECT id FROM posts ORDER BY RANDOM() LIMIT 1), (SELECT id FROM comments ORDER BY RANDOM() LIMIT 1), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (ROUND(RANDOM() * 4) + 1, (SELECT id FROM users ORDER BY RANDOM() LIMIT 1), (SELECT id FROM posts ORDER BY RANDOM() LIMIT 1), NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (ROUND(RANDOM() * 4) + 1, (SELECT id FROM users ORDER BY RANDOM() LIMIT 1), (SELECT id FROM posts ORDER BY RANDOM() LIMIT 1), NULL, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
