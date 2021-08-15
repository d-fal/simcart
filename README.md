[![Build Status](https://cloud.drone.io/api/badges/d-fal/simcart/status.svg)](https://cloud.drone.io/d-fal/simcart)
# Simple Cart Project 

A very simple cart handling project with **grpc-gateway** , **redisearch** and **postgresql**

## Code strucure:
This code is structured to be in compliance with **domain driven design (DDD)**. There are four layers that compose the code. 
- Application 
  - Application layer comprises the application skeleton or scaffold. This layer can connect to all the other layers when necessary.
- Interface
  - This layer provides a way for the clients outside to get connected to simcart.
- Domain
  - Domain layer composes all the logic inside the app. In the **simcart** we are having three subdomains, product, cart and search. All the relevant operations are included in the domain layer.
- Infrastructure
  - This layer encompasses storages connections, message brokers, etc.  

## Searching:
**Simcart** benefits from a very sophisticated **redisearch** for search purposes. As a sample setup, we have indexed *title* field and added several fields to be searched later. Below, the used schema in the code is explained.
```
	schema := redisearch.NewSchema(
		redisearch.DefaultOptions).AddField(redisearch.NewTextField("color")).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 3, Sortable: true})).
		AddField(redisearch.NewNumericField("price")).
		AddField(redisearch.NewTextField("weight")).
		AddField(redisearch.NewTextField("size")).
		AddField(redisearch.NewTextField("cat"))

```
Redisearch provides several stunning methods such as phrase search, autocompletion and suggestion based on Levenshtein distance algorithm. In this project, we would store the products and its detail at project startup and a **payload** would be added there.


### API

* Search Products
  - Searches inside the products and finds them based on the indexed fields.
* Add product: adds a product to the cart.
* Remove: removes the product of of the cart.
* Checkout: changes the satus of currently open cart and let the user to create another new cart.



### health check and metrics 
There are two types of metrics for this app. Grpc metrics and rest metrics.
* grpc health is implemented in compliance with [GRPC HEALTH CHECK PROTOCOL](https://github.com/grpc/grpc/blob/master/doc/health-checking.md).


### tracing:
Opentracing based on jaeger is implemented in this project and all the request headers are being logged.


### Protocolbuffers

This project is working based on protocol buffers. All the messages should be defined before setting up the interface. Here is an example of **protobuf** file inside **pb** folder.

```

syntax = "proto3";

package simcart.api.product.search;

option go_package = "simcart/api/pb/productpb/searchpb;searchpb";
import "google/api/annotations.proto";
import "productpb/searchpb/messages.proto";

service Search {
    rpc SearchProduct (Request) returns (Response) {
      option (google.api.http) = {
			post: "/v1/product/search"
			body: "*"
		};

    }
}

```
As you can see, **Request** and **Response** messages should be defined. [See pb/proto/](https://github.com/d-fal/simcart/tree/master/pb)
After creating the proper protocolbuffers, you should execute the folloeing command:

```
  make proto
```
### Building and running the project:
To run this project, simply run the docker-compose.yml.
```
  docker-compose up
```
This should create all the necessary apis and swagger files in [api](https://github.com/d-fal/simcart/tree/master/api) folder.
In order to use test_db, one should change the **run.sh** file a bit and run it in this way:
```
  POSTGRESS_DATABASE=test_db simcart mode --debug=true
```

### Grpc interface
Once used in debug mode, the reflection is being enabled on grpc server. Thus, you can call it by tools like **evans** and see inside it.



### Demo

You can see a demo app [here](http://simcart.100g.ir/).
You can also test the apis at [simcart.100g.ir](simcart.100g.ir) at port 80 , 50052 for rest and grpc protocols respectively.

* metrics : [GET] http://simcart.100g.ir/v1/apm/metrics
* health : [GET] http://simcart.100g.ir/v1/apm/health
* add : [POST] http://simcart.100g.ir/v1/cart/add
* list : [POST] http://simcart.100g.ir/v1/cart/list
* rm : [POST] http://simcart.100g.ir/v1/cart/rm
* checkout : [PUT] http://simcart.100g.ir/v1/cart/checkout


**grpc** interface by [**evans**](https://github.com/ktr0731/evans)

```
evans --host simcart.100g.ir -p 50052 -r 
```