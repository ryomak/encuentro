User.seed_once(:id) do |s|
  s.name = 'Admin'
  s.sex = 'male'
  s.email = 'admin@user.com'
  s.password = 'admin'
  s.password_confirmation = 'admin'
  s.birthday = '1996/07/16'
  s.status = true
  s.university = '同志社大学生'
end
