# mini-wallet-exercise

```
in this project i use mvc architecture, no database connection required for you to run this project
because it uses memDB, you can use customer_xid below inorder for you to test the project
```
customer_xid
```
ea0212d3-abd6-406f-8c67-868e814a2436
```
link : postman `https://documenter.getpostman.com/view/8411283/SVfMSqA3?version=latest#99bca41f-ecf6-4dee-a44d-154d2f8f4096`

for the route that uses token i use jwt, so the format is Authorization: Bearer
`example`
```
curl --location --request POST 'localhost:8080/api/v1/wallet' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjdXN0b21lcl94aWQiOiJlYTAyMTJkMy1hYmQ2LTQwNmYtOGM2Ny04NjhlODE0YTI0MzYiLCJleHAiOjE2NDU3NzA2OTh9.6DXs8QHCdflh3g_KxOXEjGS3fg2X4WSKphGQxChh8Vk'
```
