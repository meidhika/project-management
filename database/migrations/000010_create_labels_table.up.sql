CREATE TABLE labels(
internal_id BIGSERIAL PRIMARY KEY,
public_id UUID NOT NULL DEFAULT gen_random_uuid(),
name VARCHAR(255) NOT NULL,
created_at TIMESTAMP NOT NULL DEFAULT NOW(),
color VARCHAR(50) NOT NULL,
CONSTRAINT labels_public_id_unique UNIQUE(public_id)
);