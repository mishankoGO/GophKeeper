CREATE TABLE IF NOT EXISTS credentials (
         user_id UUID DEFAULT uuid_generate_v4(),
         login VARCHAR(50) NOT NULL UNIQUE,
         hash_password TEXT NOT NULL,
         PRIMARY KEY(user_id)
);

CREATE TABLE IF NOT EXISTS users (
        user_id UUID,
        created_at TIMESTAMP,
        PRIMARY KEY(user_id),
        CONSTRAINT fk_credential
            FOREIGN KEY(user_id)
                REFERENCES credentials(user_id)
);

CREATE TABLE IF NOT EXISTS log_passes (
        log_pass_id INT GENERATED ALWAYS AS IDENTITY,
        user_id UUID,
        name VARCHAR(50) NOT NULL,
        hash_login TEXT,
        hash_password TEXT,
        updated_at TIMESTAMP,
        meta JSONB,
        PRIMARY KEY(log_pass_id),
        CONSTRAINT fk_log_pass
            FOREIGN KEY(user_id)
                REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS texts (
        text_id INT GENERATED ALWAYS AS IDENTITY,
        user_id UUID,
        name VARCHAR(50) NOT NULL,
        hash_text TEXT,
        updated_at TIMESTAMP,
        meta JSONB,
        PRIMARY KEY(text_id),
        CONSTRAINT fk_text
            FOREIGN KEY(user_id)
                REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS cards (
        card_id INT GENERATED ALWAYS AS IDENTITY,
        user_id UUID,
        name VARCHAR(50) NOT NULL,
        hash_card_number TEXT NOT NULL,
        hash_card_holder TEXT,
        expiry_date TIMESTAMP NOT NULL,
        hash_cvv TEXT NOT NULL,
        updated_at TIMESTAMP,
        meta JSONB,
        PRIMARY KEY(card_id),
        CONSTRAINT fk_card
            FOREIGN KEY(user_id)
                REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS binary_files (
        file_id INT GENERATED ALWAYS AS IDENTITY,
        user_id UUID,
        name VARCHAR(50) NOT NULL,
        hash_file TEXT NOT NULL,
        updated_at TIMESTAMP,
        meta JSONB,
        PRIMARY KEY(file_id),
        CONSTRAINT fk_file
            FOREIGN KEY(user_id)
                REFERENCES users(user_id)
);