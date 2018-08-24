OUTPUT_DIR = $(PWD)/product
SERVER_OUTPUT_DIR = $(OUTPUT_DIR)/server
CLIENT_OUTPUT_DIR = $(OUTPUT_DIR)/client
SERVER_USER = root
SERVER_ADDR = app.juksl.com

reset-db:
	sh ./script/reset_db.sh 

deps-client:
	cd ksk && make deps
	cd ryoma && make deps

build-client: build-client-ksk build-client-ryoma

build-client-ksk:
	rm -rf $(CLIENT_OUTPUT_DIR)/ksk
	mkdir -p $(CLIENT_OUTPUT_DIR)
	cd ksk && make build
	mv ksk/dist $(CLIENT_OUTPUT_DIR)/ksk

build-client-ryoma:
	rm -rf $(CLIENT_OUTPUT_DIR)/ryoma
	mkdir -p $(CLIENT_OUTPUT_DIR)
	cd ryoma && make build
	mv ryoma/dist $(CLIENT_OUTPUT_DIR)/ryoma

develop--deps:

