-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS News (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS NewsCategories (
    news_id INT REFERENCES News(id),
    category_id INT,
    PRIMARY KEY (news_id, category_id)
);

CREATE TABLE IF NOT EXISTS Categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS NewsCategories;
DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS News;
-- +goose StatementEnd