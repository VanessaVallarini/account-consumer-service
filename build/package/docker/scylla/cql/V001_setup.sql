CREATE KEYSPACE account WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE account;
   
CREATE TABLE user (id uuid, name text, email text, PRIMARY KEY(id));

CREATE TABLE address (id uuid, alias text, city text, district text, public_place text, zip_code text, user_id text, PRIMARY KEY(id));

CREATE TABLE phone (id uuid, country_code text, area_code text, number text, user_id text, PRIMARY KEY(id));
