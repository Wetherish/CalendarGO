mysql.server start
go work init
go work use main/ Controller/ Models/ View/ Server/ data-access/ env/ tests/
cd Controller 
go get .
cd ..
cd Server
go get .
cd ..