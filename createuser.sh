mysql -h 127.0.0.1 -P 53306 -u root -pmysqlvotmdnjem <<MYSQL_SCRIPT
CREATE DATABASE $1 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
MYSQL_SCRIPT

mysql -h 127.0.0.1 -P 53306 -u root -pmysqlvotmdnjem <<MYSQL_SCRIPT
CREATE USER '$2'@'%' IDENTIFIED BY '$3';
GRANT ALL PRIVILEGES ON $1.* TO "$2"@'%';
FLUSH PRIVILEGES;
MYSQL_SCRIPT

echo "Username: $2 on $1"
echo "Password: [$3]"

