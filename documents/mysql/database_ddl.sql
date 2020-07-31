-- executed by root user
SELECT version();   -- 8.0.21
CREATE DATABASE IF NOT EXISTS utilsys;
CREATE USER 'app_be'@'localhost';
ALTER USER 'app_be'@'localhost' IDENTIFIED BY 'hogehoge';
GRANT ALL PRIVILEGES ON utilsys.* TO 'app_be'@'localhost';
SELECT host, user FROM mysql.user;
SHOW GRANTS FOR 'app_be'@'localhost';
