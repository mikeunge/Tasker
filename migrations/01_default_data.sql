/* INSERT default values */
INSERT INTO projects ("id", "name") VALUES (0, "default");
INSERT INTO status ("title") VALUES ("todo");
INSERT INTO status ("title") VALUES ("in progress");
INSERT INTO status ("title") VALUES ("open questions");
INSERT INTO status ("title") VALUES ("do testing");
INSERT INTO status ("title") VALUES ("peer review");
INSERT INTO status ("title") VALUES ("qa");
INSERT INTO status ("title") VALUES ("done");

/* INSERT set WHAT migration is currently deployed */
INSERT INTO logs ("code", "action", "caller", "message") VALUES ("MIGRATION", "created", "SYSTEM", "Ran migration: 01_default_data.sql");