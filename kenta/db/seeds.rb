# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)

ActiveRecord::Base.transaction do
  #Userサンプル
  10.times do |i|
    User.find_or_initialize_by({ name: "松岡裕典#{i}" }).update!({ name: "松岡裕典#{i}", sex: "female", email: "sample#{i}@sample.com", birthday: "1996/1/05", university: "同志社", password: "sample#{i}", password_confirmation: "sample#{i}"});
  end

  #Planサンプル
   5.times do |i|
     user = User.find(i+1)
     plan = user.plans.find_or_initialize_by({ date: "2018/10/#{i+1}" })
     plan.update!({ date: "2018/10/#{i+1}", place: "ラブホテル", drink: "あり", course: "おっぱいコース" });
   end
end