OUTPUT_DIR = $(PWD)/product
SERVER_OUTPUT_DIR = $(OUTPUT_DIR)/server
CLIENT_OUTPUT_DIR = $(OUTPUT_DIR)/client
SERVER_USER = root
SERVER_ADDR = app.juksl.com

reset-db:
	sh ./script/reset_db.sh 
  
develop-push:
	git subtree push --prefix=client heroku master
