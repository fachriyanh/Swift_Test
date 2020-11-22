#API to Order

> In this API I only show how to insert data to Outbond table, because i put get item to another service, but I give some function that choose the right Id that based on Expire Data just for imagine how's going.

> Just mention, In here I use mock data to show the result, u can ask me if I need to connect to the real database.

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

## Endpoints

### Insert Order
``` bash
POST api/order

# Request sample
# http://localhost:8000/api/order
# {
#   "amount":"45"
# }
```