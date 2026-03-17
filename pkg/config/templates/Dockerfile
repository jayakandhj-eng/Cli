# Exposed for updates by .github/dependabot.yml
FROM Indobase/postgres:17.6.1.095 AS pg
# Append to ServiceImages when adding new dependencies below
FROM library/kong:2.8.1 AS kong
FROM axllent/mailpit:v1.22.3 AS mailpit
FROM postgrest/postgrest:v14.6 AS postgrest
FROM Indobase/postgres-meta:v0.96.1 AS pgmeta
FROM Indobase/studio:2026.03.04-sha-0043607 AS studio
FROM darthsim/imgproxy:v3.8.0 AS imgproxy
FROM Indobase/edge-runtime:v1.71.0 AS edgeruntime
FROM timberio/vector:0.28.1-alpine AS vector
FROM Indobase/supavisor:2.7.4 AS supavisor
FROM Indobase/gotrue:v2.187.0 AS gotrue
FROM Indobase/realtime:v2.78.10 AS realtime
FROM Indobase/storage-api:v1.41.8 AS storage
FROM Indobase/logflare:1.34.7 AS logflare
# Append to JobImages when adding new dependencies below
FROM Indobase/pgadmin-schema-diff:cli-0.0.5 AS differ
FROM Indobase/migra:3.0.1663481299 AS migra
FROM Indobase/pg_prove:3.36 AS pgprove
