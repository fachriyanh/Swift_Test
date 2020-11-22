Because excel is an old technology for data processing, so, the features are classified as traditional. In my opinion, the data structure can be simplified by connecting the data which are interconnected. For the design system, in my opinion, this is the best design that I made:

![design](https://user-images.githubusercontent.com/43384324/99877867-c2e9f980-2c33-11eb-9e24-4e9f1c0d1bda.PNG)

The product table is used for master data. I assume that the assets in and out of a warehouse are always the same type. Therefore the SKU, Name and Expirable never change. So that if there is Inbound activity, it will first refer to this table. If in a special case, there is a new product, then this product table should be added.

In the inbound_table contains every product that enters the warehouse. I assumed that each type of product that entered simultaneously has the same details including the expiration date. SKU Number is a foreign key that refer to product_table.

stock_table table shows stock availability in a warehouse. Stock_table grouping is not based on product only, but also the details in the inbound table. The aims are to differentiate between products that have an expiry date. so, they will be taken out by this order.

outbound_table is a table to record what items leave the warehouse. This happens when the user has completed the payment. Items that have been selected based on the expiry date or inbound date, will be inserted into this table. This table requires a product_id that affects stock availability in the warehouse. This table also provides status attributes that are linked to the status table which is useful for checking delivery status.


#Services :

#show product
to insert new item as an inbound, we need to know what product that support with this system (based on the assumption before).
parameter : null
return : name, sku

#add inbound
when the warehouse got a new stock of some items, the transaction log, will record by inbound_table. simultaneously, stock_table also update with composite primery key from sku and idinbound, and current stock as a 'jumlah' inbound.
parameter : sku, expirydate, jumlah, harga, total, noOrder
return : success -> jumlah, error -> error message

#show available product
we also need to monitor all product that available in the warehouse. we can use stock_table to realize this service that related to inbound_table that has more specific data.
parameter : null
return : success -> name, currentstock, expirydate

#get item outbond
this service called when customer do booking and the status is waiting for the payment. this service used to check the Expirable. If the product has the expiration date, then it will be based on the expiration date. But if the product does not have an expiration date, it will be based on the inbound date.
parameter : expirable
return : idInbound

#add outbond
this service use to record the products going out of inventory when there is an order. first, we need getItemOutbond table to find idOutbond. then we insert it with another data to outbond_table
parameter : idProduct, quantity, price
return : total, status

#tracking (order status)
this status will tell us the order that outbond from warehouse
parameter : idProduct
return : status


Based on the table I created that contains a lot of relations, so I chose SQL database to make it happen. PostgreSQL and MySQL would fit this concept. For its service, a technology that is light and fast in execution can be relied on. Like Golang for example, which is the latest technology and is being massively used. The concurrency feature of Golang also helps with precise execution in units of time.