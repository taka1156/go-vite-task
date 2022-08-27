/* ロール */
INSERT INTO roles (role_icon, role_name, created_at, updated_at, deleted_at)
VALUES
    (NULL, "一般ユーザー",  "2022-08-26 09:00:00", NULL, NULL),
    (NULL, "管理者",       "2022-08-26 10:00:00", NULL, NULL),
    (NULL, "タスク管理担当","2022-08-26 11:00:00", NULL, NULL),
    (NULL, "開発者",       "2022-08-26 11:00:00", NULL, NULL);


/* ユーザー */
/* TODO: パスワードを平文で登録することになるので後日、スクリプトから作成するように変更する */
INSERT INTO users (user_icon,user_name,email, password, role_id, created_at, updated_at, deleted_at)
VALUES
    (NULL, "鈴木", "example1@exmple.com", "password1", 1, "2022-08-27 09:00:00", NULL, NULL), /* go */
    (NULL, "田中", "example2@exmple.com", "password2", 1, "2022-08-27 09:00:00", NULL, NULL), /* go */
    (NULL, "小川", "example3@exmple.com", "password3", 2, "2022-08-27 09:00:00", NULL, NULL), /* JS */
    (NULL, "小島", "example4@exmple.com", "password4", 3, "2022-08-27 09:00:00", NULL, NULL), /* JS */
    (NULL, "松永", "example5@exmple.com", "password4", 4, "2022-08-27 09:00:00", NULL, NULL); /* PHP */

/*  チーム */
INSERT INTO teams (team_name, team_icon, team_memo, start_date, end_date, is_end, created_at, updated_at, deleted_at)
VALUES
    ("A", NULL, "AAA", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL),
    ("B", NULL, "BBB", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL),
    ("C", NULL, "CCC", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL),
    ("D", NULL, "DDD", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL),
    ("E", NULL, "EEE", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL),
    ("F", NULL, "FFF", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL),
    ("G", NULL, "GGG", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL),
    ("H", NULL, "HHH", "2022-08-26 09:00:00", NULL, false, "2022-08-27 09:00:00", NULL, NULL);

/* カテゴリー */
INSERT INTO categories (category_name, category_icon, created_at, updated_at, deleted_at)
VALUES
    ("HTML",               NULL, "2022-08-26 09:00:00", NULL, NULL),
    ("CSS",                NULL,"2022-08-27 09:00:00", NULL, NULL),
    ("SCSS/SASS",          NULL,"2022-08-27 09:00:00", NULL, NULL),
    ("JavaScript",         NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("TypeScript",         NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("Golang",             NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("PHP",                NULL,"2022-08-27 09:00:00", NULL, NULL),
    ("Ruby",               NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("Python",             NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("SQL",                NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("ALL_LANGAGE",        NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("インフラ",            NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("ヒューマンスキル",      NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("ビジネス",            NULL, "2022-08-27 09:00:00", NULL, NULL),
    ("エンジニアリングスキル", NULL, "2022-08-27 09:00:00", NULL, NULL);

/* タスク */
INSERT INTO tasks (task_name, link, level, description, is_review, is_all_required, category_id, created_at, updated_at, deleted_at)
VALUES
    ("golangの基礎文法",                   "exmaple.com",  3, "## fmtでコンソール出力", false, false, 6, "2022-08-27 09:00:00", NULL, NULL),
    ("Dockerの使い方",                    "exmaple.com",   4, "## 環境構築",          true, true,  12,  "2022-08-27 09:00:00", NULL, NULL),
    ("Reactの使い方",                     "exmaple.com",   5, "## CRAの導入及び準備",  true, false,  4,  "2022-08-27 09:00:00", NULL, NULL),
    ("Laravelの環境でTODO管理サイトを作る", "exmaple.com",    6, "## モデルの作成",       true, false,  7,  "2022-08-27 09:00:00", NULL, NULL), 
    ("Express.jsを利用してCRUD処理作成",   "exmaple.com",    6, "## MySQLとの接続",     true, false,  4, "2022-08-27 09:00:00", NULL, NULL), 
    ("SQL基礎構文",                      "exmaple.com",    4, "## 外部キー制約について", true, true,  10, "2022-08-27 09:00:00", NULL, NULL);

/* タスクフラグ */
INSERT INTO task_flags (task_flag_name, created_at, updated_at, deleted_at)
VALUES
    ("未着手",    "2022-08-27 09:00:00", NULL, NULL),
    ("作業中",    "2022-08-27 09:00:00", NULL, NULL),
    ("レビュー中", "2022-08-27 09:00:00", NULL, NULL),
    ("完了",      "2022-08-27 09:00:00", NULL, NULL);

INSERT INTO task_status (task_id, user_id, task_flag_id, progress, created_at, updated_at, deleted_at)
VALUES
    (1, 1, 2, 35, "2022-08-27 09:00:00", NULL, NULL),
    (1, 2, 3, 50, "2022-08-27 09:00:00", NULL, NULL),
    (2, 1, 2, 25, "2022-08-27 09:00:00", NULL, NULL),
    (2, 2, 2, 30, "2022-08-27 09:00:00", NULL, NULL),
    (2, 3, 2, 20, "2022-08-27 09:00:00", NULL, NULL),
    (2, 4, 2, 20, "2022-08-27 09:00:00", NULL, NULL),
    (2, 5, 2, 40, "2022-08-27 09:00:00", NULL, NULL),
    (3, 3, 2, 60, "2022-08-27 09:00:00", NULL, NULL),
    (3, 4, 3, 75, "2022-08-27 09:00:00", NULL, NULL),
    (4, 5, 2, 50, "2022-08-27 09:00:00", NULL, NULL),
    (5, 3, 3, 80, "2022-08-27 09:00:00", NULL, NULL),
    (5, 4, 3, 80, "2022-08-27 09:00:00", NULL, NULL),
    (6, 1, 2, 10, "2022-08-27 09:00:00", NULL, NULL),
    (6, 2, 1,  0, "2022-08-27 09:00:00", NULL, NULL),
    (6, 3, 2, 30, "2022-08-27 09:00:00", NULL, NULL),
    (6, 4, 1,  0, "2022-08-27 09:00:00", NULL, NULL),
    (6, 5, 2, 50, "2022-08-27 09:00:00", NULL, NULL);
