class User < ApplicationRecord
  has_secure_password
  has_many :plans, dependent: :destroy
  validates :name, presence: true
  validates :sex, presence: true
  validates :email, presence: true, uniqueness: true
  validates :birthday, presence: true
end
