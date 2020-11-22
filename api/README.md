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

![order](https://user-images.githubusercontent.com/43384324/99898143-e5325480-2cd1-11eb-9a18-07e63d76ce2a.PNG)