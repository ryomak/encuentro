class CreateUsers < ActiveRecord::Migration[5.1]
  def change
    create_table :users do |t|
      t.string :name, null: false
      t.string :sex, null: false
      t.string :email, null: false, index: { unique: true }
      t.string :birthday, null: false
      t.string :university

      t.timestamps
    end
  end
end
