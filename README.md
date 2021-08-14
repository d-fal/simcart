# Simple Cart Project 

A very simple cart handling project with **grpc-gateway** , **redisearch** and **postgresql**


### API

* Search Products
  - Searches inside the products and finds them based on the indexed fields.
* Add product: adds a product to the cart.
* Remove: removes the product of of the cart.
* Checkout: changes the satus of currently open cart and let the user to create another new cart.

### health check and metrics 
There are two types of metrics for this app. Grpc metrics and rest metrics.
* grpc health is implemented in compliance with [GRPC HEALTH CHECK PROTOCOL](https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

### Demo

You can see a demo app [here](http://simcart.100g.ir/).