# Ecom Microservice

## Functional Requirements
1) User should be able to see all products
2) User should be able to register and login (No sso login)
3) User should be able to add product to cart
4) User should be able to checkout and make payment and place order
5) User should be able to see his orders list
6) Admin should be able to create product
7) Admin should be able to delete product

## Decomposition the ecom domain into microservices based on:
1) Business Capability
Products,orders,Customers,Payments, Admin
2) Based on subdomain
Products,orders,customers,payments,Cart,Admin

Communication between microservices
1) Rest Sync (Request response model) between client and microservices. 
This will be later moved to APIGW. 
2) Maybe GRAPHQL
3) GRPC between microservices
4) Async wherever possible

More detail regarding the api design in each microservice 



[buf.yaml](product%2Fbuf.yaml)