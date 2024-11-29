DROP TABLE IF EXISTS todos;
CREATE TABLE todos(
  id          uuid PRIMARY KEY,
  content     text NOT NULL,
  completed   boolean NOT NULL 
);