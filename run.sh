#!/usr/bin/env sh

./simcart migrate create model --extensions=uuid-ossp,hstore
./simcart seed
./simcart
