class CreateUsers < ActiveRecord::Migration[5.1]
  def change
    create_table :users do |t|
      t.string :name, null: false
      t.string :sex, null: false
      t.string :password_digest, null: false
      t.string :email, null: false, index: { unique: true }
      t.datetime :birthday, null: false
      t.string :university
      t.boolean :status, null: false, default: false

      t.timestamps
    end
  end
end
