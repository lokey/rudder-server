
--
-- wh_table_uploads
--

ALTER TABLE wh_uploads ADD COLUMN IF NOT EXISTS metadata JSONB;
