CREATE USER IF NOT EXISTS 'app'@'localhost' ;
drop database if exists encuentro;
create database encuentro;
grant all on encuentro.* to app;
