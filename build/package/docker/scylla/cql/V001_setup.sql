CREATE KEYSPACE account WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE account;

CREATE TABLE address (id uuid,alias text,city text,district text ,public_place text ,zip_code text, PRIMARY KEY(id));

CREATE TABLE phone (id uuid,country_code text,are_code text,number text, PRIMARY KEY(id));

CREATE TABLE user (id uuid,address_id text, phone_id text, name text,email text, PRIMARY KEY(id, address_id, phone_id));