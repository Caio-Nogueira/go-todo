-- Inserting data into the "users" table

-- Inserting data into the "users" table
INSERT INTO users (username, password, email) VALUES ('johndoe', '$2a$10$Jq/VOw6H0hxESwDzSZzsZenwhLutlzgzUlndMoL46eZeJfSl88.Dq', 'johndoe@email.com');
INSERT INTO users (username, password, email) VALUES ('janedoe', '$2a$10$50luXhZb10anSJduZeJoOOdlP5BcqF/JUufNvb3kYUpdHHV7Zgmlq', 'janedoe@email.com');


INSERT INTO todo (title, description, user_id) VALUES ('Buy groceries', 'Milk, bread, eggs, and cheese', (SELECT id FROM users WHERE username = 'johndoe'));
INSERT INTO todo (title, description, user_id) VALUES ('Finish project', 'Complete the final report and submit it', (SELECT id FROM users WHERE username = 'johndoe'));
INSERT INTO todo (title, description, user_id) VALUES ('Call mom', 'Ask her how she is doing and catch up', (SELECT id FROM users WHERE username = 'janedoe'));
INSERT INTO todo (title, description, user_id) VALUES ('Go for a run', 'Jog for 30 minutes around the park', (SELECT id FROM users WHERE username = 'janedoe'));
INSERT INTO todo (title, description, user_id) VALUES ('Read a book', 'Spend an hour reading a novel', (SELECT id FROM users WHERE username = 'johndoe'));
INSERT INTO todo (title, description, user_id) VALUES ('Write a blog post', 'Write a 500-word blog post about a recent trip', (SELECT id FROM users WHERE username = 'johndoe'));
INSERT INTO todo (title, description, user_id) VALUES ('Do laundry', 'Wash and fold clothes', (SELECT id FROM users WHERE username = 'janedoe'));
INSERT INTO todo (title, description, user_id) VALUES ('Clean the house', 'Vacuum, dust, and mop the floors', (SELECT id FROM users WHERE username = 'janedoe'));
