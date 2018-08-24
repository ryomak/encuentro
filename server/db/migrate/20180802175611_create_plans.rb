class CreatePlans < ActiveRecord::Migration[5.1]
  def change
    create_table :plans do |t|
      t.belongs_to :user, null: false, foreign_key: true
      t.datetime :date, null: false
      t.string :place, null: false
      t.string :drink, null: false
      t.string :course, null: false
      t.integer :status, null: false, default: 0

      t.timestamps
    end
  end
end
