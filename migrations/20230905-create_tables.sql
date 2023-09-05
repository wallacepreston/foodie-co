-- run with psql -U postgres -d foodie -a -f migrations/20230905-create_tables.sql

-- Create the "recipes" table
DROP TABLE IF EXISTS recipes;
CREATE TABLE recipes (
    recipe_id serial PRIMARY KEY,
    name text,
    instructions text,
    created_at date,
    updated_at date,
    deleted_at date
);

-- Insert sample data
INSERT INTO recipes (name, instructions, created_at, updated_at)
VALUES
    ('Spicy Thai Red Curry', 'Cooking instructions: Mix red curry paste, coconut milk, and your choice of protein for a delicious Thai dish.', '2023-09-05', '2023-09-05'),
    ('Mouthwatering Margherita Pizza', 'Instructions: Roll out pizza dough, add tomato sauce, mozzarella cheese, fresh basil, and bake until golden brown.', '2023-09-05', '2023-09-05'),
    ('Chocolate Lovers Brownies', 'Baking Directions: Prepare brownie batter with rich dark chocolate. Bake until the top is cracked and the edges are set.', '2023-09-05', '2023-09-05'),
    ('Savory Veggie Stir-Fry', 'Stir-Fry Steps: Sauté your favorite vegetables with garlic, ginger, and soy sauce for a quick and healthy meal.', '2023-09-05', '2023-09-05'),
    ('Deluxe Breakfast Burrito', 'Cooking Tips: Scramble eggs, add cooked bacon, cheese, and salsa. Roll into a tortilla for a hearty breakfast burrito.', '2023-09-05', '2023-09-05');