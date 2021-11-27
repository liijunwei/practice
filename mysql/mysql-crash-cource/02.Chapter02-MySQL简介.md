page 8

man mysql
```sh
mysql -u xxx -p -hyyy -D db
       Then type an SQL statement, end it with ;, \g, or \G and press Enter.

mysql -u root -p -E -D some_db

# For beginners, a useful startup option is --safe-updates (or --i-am-a-dummy, which has the same effect). Safe-updates mode is helpful for cases when
mysql -u root -p -E --i-am-a-dummy -D some_db

```

```ruby
[3] pry(main)> 'a'.ord
=> 97
```

```sql
mysql> select char(97);
+----------+
| char(97) |
+----------+
| a        |
+----------+
1 row in set (0.00 sec)
```

page 16