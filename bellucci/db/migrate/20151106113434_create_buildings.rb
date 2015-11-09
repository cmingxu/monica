class CreateBuildings < ActiveRecord::Migration
  def change
    create_table :buildings do |t|
      t.string :name
      t.integer :level
      t.integer :city_id
      t.integer :center_x
      t.integer :center_y
      t.integer :height
      t.integer :weight

      t.timestamps null: false
    end
  end
end
