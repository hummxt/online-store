This is an online shop backend I built with using the Go language. The Gin framework is used to manage web traffic. Product and user data is stored in a PostgreSQL database, with GORM facilitating communication with the database. Redis is used to cache frequently requested information for faster response times. Security is implemented using tokens to authenticate users and protect data.

Project Features:
- User signup and login system.
- Profile section for updating email, username, name, phone number, and address.
- Product display and search functionality.
- Product categorization.
- Shopping cart for managing purchases.
- Checkout system for converting carts into real orders.

Detailed information about API endpoints, including methods, URLs, and required fields, can be found in the [APIs.md](./APIs.md) file.
