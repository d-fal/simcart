#!/usr/bin/env sh

./simcart create model --extensions=uuid-ossp,hstore
./simcart seed
./simcart
