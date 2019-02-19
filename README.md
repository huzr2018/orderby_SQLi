Kohana/koseven allows SQL Injection via the order_by parameter.
1、In some website, developers use dynamic sort operation like this
`http://x.com?order_by=desc`
attacker can execute SQL like:
```
http://x.com?order_by=,(extractvalue(1,concat(0x3a,substring((select group_concat(column_name) from information_schema.columns where table_name like 'xxx'),1,30) )))
```
or other way like:
http://www.securityidiots.com/Web-Pentest/SQL-Injection/group-by-and-order-by-sql-injection.html
2、If we don't remind developers the potential risks of using untrusted input at order_by ,sql injection may occur.
3、developers may think the instructions are safe and sure of security. Most of them think the use of orm is 100% guard against sql injection.
4、So can we add a fun like 'order_by_safe' which receive field only in models