CREATE DATABASE devbook;

CREATE USER golang WITH PASSWORD 'golang';

\c devbook;

-- DROP TABLE IF EXISTS posts;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS followers;

CREATE TABLE users(
    id SERIAL PRIMARY KEY, 
    name varchar(50) not null,
    username varchar(50) unique not null,
    email varchar(50) unique not null,
    password varchar(100) not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE users IS 'Users data';
COMMENT ON COLUMN users.id IS 'Tuple primary key';
COMMENT ON COLUMN users.name IS 'User name';
COMMENT ON COLUMN users.email IS 'User email';
COMMENT ON COLUMN users.created_at IS 'created time';
COMMENT ON COLUMN users.updated_at IS 'updated time';
COMMENT ON COLUMN users.deleted_at IS 'deleted time';

INSERT INTO users (name, username, email, password) VALUES
    ('Usuário 1', 'usuario_1', 'usuario1@gmail.com', '$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy'),
    ('Usuário 2', 'usuario_2', 'usuario2@gmail.com', '$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy'),
    ('Usuário 3', 'usuario_3', 'usuario3@gmail.com', '$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy');

CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    PRIMARY KEY (user_id, follower_id),

    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE followers IS 'Followers data';
COMMENT ON COLUMN followers.user_id IS 'Followed user';
COMMENT ON COLUMN followers.follower_id IS 'Follower user';
COMMENT ON COLUMN followers.created_at IS 'created time';
COMMENT ON COLUMN followers.updated_at IS 'updated time';
COMMENT ON COLUMN followers.deleted_at IS 'deleted time';

INSERT INTO followers (user_id, follower_id) VALUES
    (1, 2),
    (3, 1),
    (1, 3);

CREATE TABLE posts(
    id SERIAL PRIMARY KEY, 
    title varchar(50) not null,
    content varchar(300) not null,
    author_id int not null,
    likes int DEFAULT 0,

    FOREIGN KEY (author_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO posts (title, content, author_id) VALUES
    ('Publicação 1 do Usuário 1', 'Essa é uma públicação do user 1', '1'),
    ('Publicação 1 do Usuário 2', 'usuario_2 acaba de publicar algo', '2'),
    ('Publicação 1 do Usuário 3', 'Dessa vez, foi o usuario_3 que publicou', '3'),
    ('Publicação 2 do Usuário 1', 'Essa é outra públicação do user 1', '1'),
    ('Publicação 2 do Usuário 2', 'usuario_2 acaba de publicar algo outra vez', '2'),
    ('Publicação 2 do Usuário 3', 'Dessa vez, foi o usuario_3 que publicou mais uma vez', '3');

GRANT ALL PRIVILEGES ON DATABASE devbook TO golang;
