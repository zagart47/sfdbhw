DROP TABLE IF EXISTS "users", "tasks_labels", "labels", "tasks";

CREATE TABLE "users"
(
    "id"   SERIAL,
    "name" TEXT NOT NULL
);
ALTER TABLE
    "users"
    ADD PRIMARY KEY ("id");

CREATE TABLE "tasks_labels"
(
    "task_id"  INTEGER NOT NULL,
    "label_id" INTEGER
);

CREATE TABLE "labels"
(
    "id"   SERIAL,
    "name" TEXT
);
ALTER TABLE
    "labels"
    ADD PRIMARY KEY ("id");

CREATE TABLE "tasks"
(
    "id"          SERIAL PRIMARY KEY,
    "opened"      BIGINT  NOT NULL,
    "closed"      BIGINT  NOT NULL,
    "author_id"   INTEGER NOT NULL,
    "assigned_id" INTEGER NOT NULL,
    "title"       TEXT    NOT NULL,
    "content"     TEXT    NOT NULL
);

ALTER TABLE
    "tasks"
    ADD CONSTRAINT "tasks_author_id_foreign" FOREIGN KEY ("author_id") REFERENCES "users" ("id");
ALTER TABLE
    "tasks"
    ADD CONSTRAINT "tasks_assigned_id_foreign" FOREIGN KEY ("assigned_id") REFERENCES "users" ("id");
ALTER TABLE
    "tasks_labels"
    ADD CONSTRAINT "tasks_labels_task_id_foreign" FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON DELETE CASCADE ;
ALTER TABLE
    "tasks_labels"
    ADD CONSTRAINT "tasks_labels_label_id_foreign" FOREIGN KEY ("label_id") REFERENCES "labels" ("id");


-- Заполнение таблицы users
INSERT INTO users (name)
VALUES ('Alice'),
       ('Bob'),
       ('Charlie'),
       ('David'),
       ('Eve'),
       ('Frank'),
       ('Grace'),
       ('Hank'),
       ('Ivy'),
       ('Jack');

-- Заполнение таблицы labels
INSERT INTO labels (name)
VALUES ('EMPTY'),
       ('Bug'),
       ('Feature'),
       ('Enhancement'),
       ('Documentation'),
       ('Urgent'),
       ('Low Priority'),
       ('High Priority'),
       ('Design'),
       ('Testing'),
       ('Refactoring');

-- Заполнение таблицы tasks
INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content)
VALUES (1633027200, 1633113600, 1, 2, 'Fix login issue', 'Users are unable to log in.'),
       (1633027200, 0, 2, 3, 'Add new feature', 'Implement a new feature for the dashboard.'),
       (1633027200, 1633113600, 3, 1, 'Improve performance', 'Optimize the database queries.'),
       (1633027200, 1633113600, 4, 5, 'Update documentation', 'Update the user manual.'),
       (1633027200, 0, 5, 6, 'Fix UI bug', 'The button is not clickable.'),
       (1633027200, 1633113600, 6, 7, 'Add unit tests', 'Write unit tests for the new feature.'),
       (1633027200, 0, 7, 8, 'Refactor code', 'Refactor the code for better readability.'),
       (1633027200, 1633113600, 8, 9, 'Design new layout', 'Create a new layout for the homepage.'),
       (1633027200, 0, 9, 10, 'Fix security issue', 'Patch a security vulnerability.'),
       (1633027200, 1633113600, 10, 1, 'Add API endpoint', 'Add a new API endpoint for the mobile app.'),
       (1633027200, 0, 1, 2, 'Optimize images', 'Compress images for faster loading.'),
       (1633027200, 1633113600, 2, 3, 'Fix broken link', 'A link in the footer is broken.'),
       (1633027200, 0, 3, 4, 'Add search functionality', 'Implement search functionality for the blog.'),
       (1633027200, 1633113600, 4, 5, 'Update privacy policy', 'Update the privacy policy document.'),
       (1633027200, 0, 5, 6, 'Fix form validation', 'Fix form validation for the contact page.'),
       (1633027200, 1633113600, 6, 7, 'Add analytics', 'Integrate analytics for the new feature.'),
       (1633027200, 0, 7, 8, 'Refactor CSS', 'Refactor CSS for better maintainability.'),
       (1633027200, 1633113600, 8, 9, 'Design new logo', 'Create a new logo for the company.'),
       (1633027200, 0, 9, 10, 'Fix email notifications', 'Email notifications are not being sent.'),
       (1633027200, 1633113600, 10, 1, 'Add social media links', 'Add social media links to the footer.');

-- Заполнение таблицы tasks_labels
INSERT INTO tasks_labels (task_id, label_id)
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 4),
       (5, 1),
       (6, 9),
       (7, 10),
       (8, 8),
       (9, 5),
       (10, 2),
       (11, 3),
       (12, 1),
       (13, 2),
       (14, 4),
       (15, 1),
       (16, 7),
       (17, 10),
       (18, 8),
       (19, 5),
       (20, 2);
