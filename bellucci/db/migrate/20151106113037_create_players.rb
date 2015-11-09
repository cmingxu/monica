class CreatePlayers < ActiveRecord::Migration
  def change
    create_table :players do |t|
      t.string :name
      t.string :email
      t.datetime :last_login_at
      t.string :last_login_ip
      t.integer :gem

      t.timestamps null: false
    end
  end
end
