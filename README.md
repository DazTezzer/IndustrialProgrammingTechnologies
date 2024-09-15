# Рулев Данила Олегович ЭФМО-02-24

## Схема базы данных интернет-магазина одежды

![image](https://github.com/user-attachments/assets/d4d640c9-0e9c-4715-9e67-0f7758a08dd8)














```

Table Customer {
  id integer [primary key]
  username varchar
  email varchar
  password_hash varchar
  created_at date
}

Table Products {
  id integer [primary key]
  name varchar
  image_id integer
  price money
  description_id integer
  category_id integer
}

Table ProductImage {
  id integer [primary key]
  name string
  image image
}

Table ProductCategory {
  id integer [primary key]
  name varchar
}

Table Orders {
  id integer [primary key]
  customer_id integer
  order_date date
  total_amount money
  status varchar
}

Table OrderItems {
  order_id integer
  product_id integer
  quantity integer
}

Table Favorites {
  id integer [primary key]
  customer_id integer
  product_id integer
}

Table ProductDescription {
  id integer [primary key]
  size string
  color string 
  material string
  gender string
  additional_info string
}

Table Reviews {
  id integer [primary key]
  product_id integer
  customer_id integer
  rating integer
  comment string
  created_at date
  order_id integer
}

Table Carts {
  id integer [primary key]
  customer_id integer
}

Table CartItems {
  cart_id integer
  product_id integer
  quantity integer
  added_at date
}

Ref: Products.image_id > ProductImage.id

Ref: Reviews.order_id > Orders.id 

Ref: CartItems.cart_id > Carts.id

Ref: Carts.customer_id > Customer.id

Ref: CartItems.product_id > Products.id

Ref: OrderItems.product_id > Products.id

Ref: OrderItems.order_id > Orders.id

Ref: Reviews.customer_id > Customer.id

Ref: Reviews.product_id > Products.id

Ref: Products.description_id > ProductDescription.id

Ref: Favorites.product_id > Products.id

Ref: Favorites.customer_id > Customer.id

Ref: Products.category_id > ProductCategory.id

Ref: Orders.customer_id > Customer.id

```
