BEGIN;
CREATE TABLE IF NOT EXISTS users.users(
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    nik VARCHAR NOT NULL,
    phonenumber VARCHAR(15),
    account_number VARCHAR(20) NOT NULL,
    latest_balance DECIMAL NOT NULL DEFAULT 0,
    created_by uuid,
    created_date TIMESTAMP,
    updated_by uuid,
    updated_date TIMESTAMP,
    deleted_by uuid,
    deleted_date TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS users_nik_phonenumber_key
ON users.users(nik, phonenumber);

CREATE UNIQUE INDEX IF NOT EXISTS users_account_number_key
ON users.users(account_number);
COMMIT;