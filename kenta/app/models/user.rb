class User < ApplicationRecord
  has_many :plans, dependent: :destroy
  validates :name, presence: true
  validates :sex, presence: true
  validates :mail, presence: true, uniqueness: true
  validates :birthday, presence: true
end
