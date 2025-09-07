-- Add index to speed up lookups by identifier
CREATE INDEX IF NOT EXISTS idx_users_identifier ON users(identifier);
