# Database Design

How each database engine actually works, and the design principles behind schema decisions.

## Oracle

- An enterprise relational database, commonly the system of record in manufacturing and large corporate environments. Strong on transactional integrity and mature tooling for large schemas.
- PL/SQL is Oracle's procedural extension to SQL, used for stored procedures and more complex logic that lives inside the database itself rather than in application code.
- Common Table Expressions (CTEs) are a clean way to structure a multi-step comparison query, for example pulling two result sets that should theoretically match, then joining them on a shared key and labeling each row as MATCH or MISMATCH, instead of nesting subqueries inside each other.

## Microsoft SQL Server

- Similar relational model to Oracle with its own dialect (T-SQL) and tooling (SQL Server Management Studio), used here alongside Oracle for reporting and data-extraction work.

## PostgreSQL

- Open-source, relational, and notably extensible. Custom types, extensions, and strong support for JSON columns when part of a schema doesn't fit a strict relational shape.
- Pairs naturally with modern typed ORMs. Prisma and Drizzle both treat it as close to a first-class target, generating fully-typed clients from the schema.

## MySQL

- The most common choice for straightforward relational workloads where the extra features of Postgres aren't needed. Simpler operationally, extremely well documented, and the default a lot of PHP-ecosystem tooling assumes.

## MongoDB

- Document-oriented instead of relational. Data is stored as JSON-like documents rather than rows across normalized tables, which fits well when the shape of a record can vary or nest naturally instead of needing to be split across many tables.

## Supabase

- A managed package built on top of Postgres: database, authentication, and realtime subscriptions bundled together, which removes a lot of the setup work a small project would otherwise spend on infrastructure before writing a single feature.
- Privileged cross-tenant queries (an admin dashboard that needs to see every tenant's data, not just the caller's own) can be handled by narrow `security definer` Postgres functions with an internal permission check, instead of introducing a service-role key into the application. No single credential in the app can bypass row-level security.

## Schema Design Principles

- Normalize to avoid duplicate/inconsistent data, but denormalize deliberately where read performance matters more than storage efficiency. The decision should be intentional either way, not a default.
- Index the columns a query actually filters or joins on; an unindexed column that's queried frequently is one of the most common causes of a database that "used to be fast."
- For a write-heavy system (like tracking frequent state changes), batching writes instead of committing on every single event reduces lock contention and keeps throughput stable as concurrent load increases. The trade-off is a small window where the very latest state hasn't been persisted yet.
