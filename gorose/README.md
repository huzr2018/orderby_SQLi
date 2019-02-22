Gorose v1.0.4 has SQL Injection when the order_by or group_by parameter can be controlled. 
   
    
then all of the three lines of code can work
```
go run dal.go 'if(1=1,id,username)'
go run dal.go 'if(1=2,id,username)'
go run dal.go 'id'
```   
In website, attacker can execute SQL like:
```
http://x.com?order_by=,(extractvalue(1,concat(0x3a,substring(
    (select group_concat(column_name) from
information_schema.columns where table_name like 'xxx'),1,30) )))
```
or other way like:   
http://www.securityidiots.com/Web-Pentest/SQL-Injection/group-by-and-order-by-sql-injection.html  
    
If we don't remind developers the potential risks of using untrusted input at order_by / group_by ,sql injection may occur.   
        
developers may think the instructions are safe and sure of security. Most of them think the use of orm is 100% guard against sql injection.   
   
So can we add a fun like 'order_by_safe' which receive field only in models.   