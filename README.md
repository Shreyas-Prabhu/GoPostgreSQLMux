# GoPostgreSQLMux
CRUD operation using Gorilla mux, Postgresql, GORM

All the apis are integrated with swagger using swaggo
swagger can be accessed using - http://localhost:4000/swagger/index.html

The table is created automatically in postgresql using Automigrate schema feature in GORM. There is no explicit table creation done. 

3 tables in database
-employees
{
  "empid":uuid.UUID(text)
  "firstname":string(text),
  "lastname":string(text),
  "gender":string(text),
  "phonenumber":int64(bigint),
  "employeeemail":string(text),
  "address":string(text),
  "bloodgroup": string(text),
  "emergencycontact":int64(bigint)
}

-assets
{
  "assetid": uuid.UUID(text),
  "assetname": string(text),
  "assettype": string(text)
}
-employee_assets
{
  "id": uuid.UUID(text),
  "empid": uuid.UUID(text),
  "assetid": uuid.UUID(text)
}

The structure of the table is same as that of model struct defined

to start application
--> go run main.go

.env file has the port and connection string to postgresql. Change the string when using your database.

Add employee body structure:
{
  "firstname":"Shreyas",
  "lastname":"Prabhu",
  "gender":"male",
  "phonenumber":8296370386,
  "employeeemail":"p@nma.com",
  "address":"Bangalore",
  "bloodgroup": "B+ve",
  "emergencycontact":7894561230
}

Add asset body structure:
{
    "assetname": "Bag",
    "assettype": "Accessories"
}

Add mapping body structure:
{
    "empid": "dc48c439-3ef8-4cc7-afb6-334ad4903dd2",
    "assetid": "3b161194-12cf-4110-85ac-0b2abda332b2"
}

NOTE: The primary keys for each table gets auto-generated


BASE URL:
http://localhost:4000

Test the api in postman/thunderclient for better view
