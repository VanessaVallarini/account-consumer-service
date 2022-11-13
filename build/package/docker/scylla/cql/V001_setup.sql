CREATE KEYSPACE account_consumer_service WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
   
CREATE TABLE account_consumer_service.account 
(
	id uuid
	,name text
	,email text
	,alias text
	,city text
	,district text
	,public_place text
	,zip_code text
	,country_code text
	,area_code text
	,number text
	,PRIMARY KEY(id)
);
