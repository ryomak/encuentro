class User < ApplicationRecord
  validates :name, presence: true
  validates :sex, presence: true
  validates :mail, presence: true, uniqueness: true
  validates :birthday, presence: true
end
