There are informations about API URLs, fields and methods.

## Auth and Profile

1. Register
Method: POST
URL: /auth/register
Fields: username, email, password, first_name, last_name

2. Login
Method: POST
URL: /auth/login
Fields: email, password
Response: token

3. Get Profile
Method: GET
URL: /profile
Note: Needs Authorization Bearer token header.

4. Update Profile
Method: PUT
URL: /profile
Fields: username, email, first_name, last_name, phone_number, address (all optional)
Note: Needs Authorization Bearer token header.

## Products

5. List Products
Method: GET
URL: /products
Query Params: page, limit

6. Product Details
Method: GET
URL: /products/:id

7. Search Products
Method: GET
URL: /products/search
Query Param: q

8. List Categories
Method: GET
URL: /categories

## Cart and Orders

9. Get Cart
Method: GET
URL: /cart

10. Add to Cart
Method: POST
URL: /cart/items
Fields: product_id, quantity

11. Update Cart Item
Method: PUT
URL: /cart/items/:id
Fields: quantity

12. Place Order
Method: POST
URL: /orders
Note: Creates order from current cart.

13. My Orders
Method: GET
URL: /orders

14. Order Details
Method: GET
URL: /orders/:id

## Admin

15. Create Product
Method: POST
URL: /admin/products
Fields: name, description, price, stock, category_id, image_urls

16. Create Category
Method: POST
URL: /admin/categories
Fields: name, description, slug
