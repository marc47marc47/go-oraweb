package statements

const (
	ORA_DATABASES = `SELECT ora_database_name FROM DUAL`

	ORA_SCHEMAS = `SELECT DISTINCT owner FROM all_objects ORDER BY owner ASC`

	ORA_INFO = `SELECT SYS_CONTEXT ('USERENV', 'SESSION_USER'),
       SYS_CONTEXT ('USERENV', 'CURRENT_USER'),
       SYS_CONTEXT ('USERENV', 'DB_NAME'),
       NULL AS current_schemas,
       SYS_CONTEXT ('USERENV', 'IP_ADDRESS'),
       NULL AS inet_client_port,
       NULL AS inet_server_addr,
       1522 AS inet_server_port,
       NULL AS version
  FROM DUAL`

	ORA_TABLE_INDEXES = `SELECT index_name, status FROM all_indexes WHERE table_name = :table_name`

	ORA_TABLE_INFO = `SELECT
  'Unk kB' AS data_size
, 'Unk kB' AS index_size
, 'Unk kB' AS total_size
, (SELECT num_rows FROM all_tables WHERE table_name = :table_name) AS rows_count`

	ORA_TABLE_SCHEMA = `SELECT table_name, column_name, data_type, nullable, data_length, character_set_name, 'not implemented' as data_default
  FROM all_tab_cols
  WHERE table_name = :table_name`

	ORA_TABLES = `SELECT table_name FROM all_tables ORDER BY owner, table_name`

	ORA_ACTIVITY = `SELECT
  null as datname,
  null as query,
  null as state,
  null as waiting,
  null as query_start,
  null as state_change,
  null as pid,
  null as datid,
  null as application_name,
  null as client_addr
  FROM dual`
)
