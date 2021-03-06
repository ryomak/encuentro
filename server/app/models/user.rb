class User < ApplicationRecord
  has_secure_password
  has_many :plans, dependent: :destroy
  validates :name, presence: true
  validates :sex, presence: true
  validates :email, presence: true, uniqueness: true
  validates :password_digest, presence: true, length: { minimum: 6 }, allow_nil: true
  validates :birthday, presence: true
  validates :status, presence: true

  enum status: { admin: true, user: false }
end
