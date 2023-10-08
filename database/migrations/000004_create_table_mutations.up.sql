BEGIN;
CREATE TABLE IF NOT EXISTS transactions.mutations(
    uuid uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_uuid uuid NOT NULL,
    trx_code varchar(5) NOT NULL,
    trx_time TIMESTAMP,
    nominal DECIMAL NOT NULL,
    created_by uuid,
    created_date TIMESTAMP,
    updated_by uuid,
    updated_date TIMESTAMP,
    deleted_by uuid,
    deleted_date TIMESTAMP
);

CREATE INDEX IF NOT EXISTS mutation_user_uuid_key
ON transactions.mutations(user_uuid);
COMMIT;