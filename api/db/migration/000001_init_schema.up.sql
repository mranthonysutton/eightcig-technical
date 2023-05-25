CREATE TABLE "employees" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "performance" bigint NOT NULL,
    "date" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX on employees (name);

INSERT INTO employees (name, performance, date)
VALUES ('John', 58, '12/11/2022');

INSERT INTO employees (name, performance, date)
VALUES ('Daniel', 87, '11/15/2022');

INSERT INTO employees (name, performance, date)
VALUES ('Sally', 34, '6/15/2022');

INSERT INTO employees (name, performance, date)
VALUES ('Tiffany', 99, '2/28/2022');
