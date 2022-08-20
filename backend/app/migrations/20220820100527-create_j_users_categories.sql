-- +migrate Up
CREATE TABLE IF NOT EXISTS j_users_categories (
    j_user_category_id INTEGER NOT NULL AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE DEFAULT NULL,
    deleted_at DATE DEFAULT NULL,
    PRIMARY KEY (j_user_category_id),
    FOREIGN KEY j_users_index1(user_id) REFERENCES users(user_id),
    FOREIGN KEY j_categories_index(category_id) REFERENCES categories(category_id)
);

-- +migrate Down
DROP TABLE IF EXISTS j_users_categories;
