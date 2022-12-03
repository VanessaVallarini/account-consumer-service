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
	,full_number text
	,PRIMARY KEY((id), email, full_number));

INSERT INTO account_consumer_service.account
(
	id
	,email
	,full_number	
	,alias	
	,city	
	,district	
	,name	
	,public_place
	,zip_code
)
VALUES
(
	'f367b7ff-caac-44e3-b869-d899da6476cf'
	,'van2@email.com'
	,'55119641282299'	
	,'SP'	
	,'São Paulo'
	,'Praça da Sé'
	,'Van'
	,'Sé'
	,'01001-000'
);

SELECT * FROM account_consumer_service.account;

--delete from account_consumer_service.account;

--DROP table account_consumer_service.account;

--select by
SELECT 
	id
	,email
	,full_number	
	,alias	
	,city	
	,district	
	,name	
	,public_place
	,zip_code
FROM account_consumer_service.account
WHERE 
email = 'van2@email.com'
OR full_number = '5511964127239'
LIMIT 1