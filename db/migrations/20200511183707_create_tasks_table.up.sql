CREATE TABLE IF NOT EXISTS tasks(
   id serial PRIMARY KEY,
   text TEXT NOT NULL,
   points SMALLINT NOT NULL,
   burnt SMALLINT NOT NULL,
   state TEXT NOT NULL
);
