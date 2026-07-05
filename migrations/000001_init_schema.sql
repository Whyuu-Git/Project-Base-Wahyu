
CREATE TABLE question_categories (
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(150) NOT NULL,
    passing_grade  INTEGER NOT NULL DEFAULT 0,
    program        VARCHAR(100),
    created_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE questions (
    id                   SERIAL PRIMARY KEY,
    text                 TEXT NOT NULL,
    number               INTEGER,
    program              VARCHAR(100),
    explanation          TEXT,
    question_category_id INTEGER NOT NULL REFERENCES question_categories(id) ON DELETE RESTRICT,
    created_at           TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_questions_category ON questions(question_category_id);


CREATE TABLE answers (
    id           SERIAL PRIMARY KEY,
    score        NUMERIC(6,2) DEFAULT 0,
    option       VARCHAR(5) NOT NULL,      -- A, B, C, D, dst
    text         TEXT NOT NULL,
    is_true      BOOLEAN NOT NULL DEFAULT FALSE,
    questions_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_answers_question ON answers(questions_id);


CREATE TABLE module (
    id         SERIAL PRIMARY KEY,
    code       VARCHAR(100) NOT NULL UNIQUE,
    name       VARCHAR(150) NOT NULL,
    program    VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE module_questions (
    id          SERIAL PRIMARY KEY,
    module_id   INTEGER NOT NULL REFERENCES module(id) ON DELETE CASCADE,
    question_id INTEGER NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (module_id, question_id)
);

CREATE INDEX idx_module_questions_module ON module_questions(module_id);
CREATE INDEX idx_module_questions_question ON module_questions(question_id);


CREATE TABLE tryout_codes (
    id          SERIAL PRIMARY KEY,
    code        VARCHAR(100) NOT NULL UNIQUE,
    name        VARCHAR(150) NOT NULL,
    start_date  TIMESTAMP NOT NULL,
    end_date    TIMESTAMP NOT NULL,
    module_id   INTEGER NOT NULL REFERENCES module(id) ON DELETE RESTRICT,
    instruction TEXT,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_tryout_codes_module ON tryout_codes(module_id);
CREATE INDEX idx_tryout_codes_daterange ON tryout_codes(start_date, end_date);


CREATE TABLE student (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(150) NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    address    TEXT,
    phone      VARCHAR(30),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE log_exam (
    id             SERIAL PRIMARY KEY,
    tryout_code_id INTEGER NOT NULL REFERENCES tryout_codes(id) ON DELETE RESTRICT,
    pass_status    BOOLEAN,
    total_score    NUMERIC(6,2) DEFAULT 0,
    repeat         INTEGER NOT NULL DEFAULT 1,
    start_date     TIMESTAMP,
    end_date       TIMESTAMP,
    student_id     INTEGER NOT NULL REFERENCES student(id) ON DELETE RESTRICT,
    created_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_log_exam_tryout ON log_exam(tryout_code_id);
CREATE INDEX idx_log_exam_student ON log_exam(student_id);


CREATE TABLE detail_log (
    id          SERIAL PRIMARY KEY,
    log_id      INTEGER NOT NULL REFERENCES log_exam(id) ON DELETE CASCADE,
    category_id INTEGER NOT NULL REFERENCES question_categories(id) ON DELETE RESTRICT,
    score       NUMERIC(6,2) DEFAULT 0,
    pass_status BOOLEAN,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_detail_log_log ON detail_log(log_id);
CREATE INDEX idx_detail_log_category ON detail_log(category_id);


CREATE TABLE history_answer (
    id                    SERIAL PRIMARY KEY,
    log_id                INTEGER NOT NULL REFERENCES log_exam(id) ON DELETE CASCADE,
    question_id           INTEGER NOT NULL REFERENCES questions(id) ON DELETE RESTRICT,
    answer_id             INTEGER,   
    number                INTEGER,   
    question              TEXT,      
    explanations          TEXT,      
    question_category_id  INTEGER NOT NULL REFERENCES question_categories(id) ON DELETE RESTRICT,
    created_at            TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_history_answer_log ON history_answer(log_id);
CREATE INDEX idx_history_answer_question ON history_answer(question_id);
CREATE INDEX idx_history_answer_category ON history_answer(question_category_id);

CREATE TABLE answer_questions (
    id                SERIAL PRIMARY KEY,
    history_answer_id INTEGER NOT NULL REFERENCES history_answer(id) ON DELETE CASCADE,
    answer_id         INTEGER REFERENCES answers(id) ON DELETE RESTRICT, -- referensi ke jawaban master asli
    option            VARCHAR(5) NOT NULL,   -- opsi versi acakan (A/B/C/D)
    text              TEXT NOT NULL,
    is_true           BOOLEAN NOT NULL DEFAULT FALSE,
    created_at        TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_answer_questions_history ON answer_questions(history_answer_id);
CREATE INDEX idx_answer_questions_answer ON answer_questions(answer_id);


ALTER TABLE history_answer
    ADD CONSTRAINT fk_history_answer_selected_answer
    FOREIGN KEY (answer_id) REFERENCES answer_questions(id) ON DELETE SET NULL;


