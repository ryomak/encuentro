CREATE USER IF NOT EXISTS 'app'@'localhost' IDENTIFIED BY 'password';
drop database IF exists encuentro;
create database encuentro;
GRANT ALL ON `encuentro`.* TO 'app'@'localhost' IDENTIFIED BY 'password';
