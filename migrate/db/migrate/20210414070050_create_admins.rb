class CreateAdmins < ActiveRecord::Migration[6.0]
  def change
    create_table :admins do |t|
      t.bigint :number
      t.string :name
      t.integer :age
      t.boolean :active
      t.text :note
      t.datetime :created_at, null: false
      t.datetime :updated_at, null: false
    end
    create_table :tasks do |t|
      t.references :admin
      t.string :name
      t.datetime :created_at, null: false
      t.datetime :updated_at, null: false
    end
    add_index :admins, :number, unique: true
  end
end
