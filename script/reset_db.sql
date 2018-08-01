drop database if exists encuentro;
create database encuentro;
grant all on encuentro.* to 'app'@'localhost' IDENTIFIED BY 'password';
