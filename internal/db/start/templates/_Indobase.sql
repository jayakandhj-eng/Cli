CREATE DATABASE _Indobase WITH OWNER postgres;

-- Switch to the newly created _Indobase database
\c _Indobase
-- Create schemas in _Indobase database for
-- internals tools and reports to not overload user database
-- with non-user activity
CREATE SCHEMA IF NOT EXISTS _analytics;
ALTER SCHEMA _analytics OWNER TO postgres;

CREATE SCHEMA IF NOT EXISTS _supavisor;
ALTER SCHEMA _supavisor OWNER TO postgres;
\c postgres

