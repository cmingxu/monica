class CreateCities < ActiveRecord::Migration
  def change
    create_table :cities do |t|
      t.string :name
      t.integer :height
      t.integer :weight
      t.integer :center_x
      t.integer :center_y
      t.boolean :is_home

      t.timestamps null: false
    end
  end
end
