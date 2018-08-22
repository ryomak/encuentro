class Plan < ApplicationRecord
  belongs_to :user
  validates :date, presence: true
  validates :place, presence: true
  validates :drink, presence: true
  validates :course, presence: true
  validates :status, presence: true
end
