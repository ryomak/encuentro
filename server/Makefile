db/init:
	rails db:create
	rails db:migrate

db/seed:
	rails db:seed

db/seed_user:
	rails db:seed_fu FILTER=create_user.rb
	rails db:seed_fu FILTER=create_admin_user.rb
