DROP TRIGGER IF EXISTS segment_deleted ON segments;

DROP FUNCTION IF EXISTS segment_deleted_trigger;

DROP TABLE IF EXISTS relations;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS segments;

DROP TABLE IF EXISTS operations;

DROP TYPE IF EXISTS session_type;

DROP TYPE IF EXISTS operation_type;