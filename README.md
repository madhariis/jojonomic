

## Postman Collection
https://documenter.getpostman.com/view/6177110/TWDdiYeH

## Endpoints
### Root List
Get list of a root of documents. The service will show all document and folder from the user's company. The company_id is provided by jwt token.
```
GET http://api-gateway.co.id/document-service
```
Response:
```
{
    "error": false,
    "data": [
         {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3d",
             "name": "Folder Tech",
             "type": "folder",
             "is_public": true,
             "owner_id": 1231,
             "share": [123,232,4333],
             "timestamp": 16576232323,
             "company_id": 23
         },
         {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
             "name": "Folder hrd",
             "type": "folder",
             "is_public": true,
             "owner_id": 1231,
             "share": [123,232,4333],
             "timestamp": 16576232323,
             "company_id": 23
         },
        {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
             "name": "Document Job desc Tech",
             "type": "document",
             "owner_id": 1231,
             "share": [123,232,4333],
             "timestamp": 16576232323,
             "company_id": 23
         }
    ]
}
```
If is_public is true, all user from a company can see the file or folder. If is_public is false, only owner and user in share can see the documents.

### Set Folder
If the ID exist, system will update the folder. If the Folder ID not exist, system will create the folder. Every folder has owner_id. You can get owner_id from json web token.
```
POST http://api-gateway.co.id/document-service/folder
```
Request Body:
```
{
        "id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3d",
        "name": "Folder Baru", 
        "timestamp": 16576232323
}
```
Response Body:
```
{
  "error": false,
  "message": "folder created",
  "data": {
      "id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3d",
      "name": "Folder Baru",
      "type": "folder",
      "content": {},
      "timestamp": 1605081795, 
      "owner_id": 123,
      "company_id": 23
  }
}
```

### Delete Folder
```
DELETE http://api-gateway.co.id/document-service/folder
```
Request Body:
```
{
        "id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3d" //id folder yang dihapus
}
```
Response Body:
```
{
  "error": false,
  "message": "Success delete folder"
}
```
### List File Per Folder ###
```
GET http://api-gateway.co.id/document-service/folder/:folder_id
```
Response Body
```
{
    "error": false,
    "data": [
        {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
             "name": "Document Job desc Tech",
             "type": "document",
             "folder_id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3e", // string kosong atau ada id foldernya
             "content": {
                 "blocks": [
                     {
                         "type": "paragraph",
                         "text": "This is paragraph"
                     }
                 ]
             }, // block editorjs
             "timestamp": 1605081795, //second unix timestamp
             "owner_id": 123, //user yg lagi login. atau user yang mau diserah terimakan dokumennya. Kalau 0, backend ambil dari data login
             "share": [1,23,4232,121] //share ke user lain
        },
        {
             "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
             "name": "Document Job desc BA",
             "type": "document"
         }
    ]
}
```
### Set (Create/Update) document ###
```
POST http://api-gateway.co.id/document-service/document
```
Request Body
```
{
    "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
    "name": "Document Job desc Tech",
    "type": "document",
    "folder_id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3e", // string kosong atau ada id foldernya
    "content": {
        "blocks": [
            {
                "type": "paragraph",
                "text": "This is paragraph"
            }
            ]
    }, // block editorjs
    "timestamp": 1605081795, //second unix timestamp
    "owner_id": 123, //user yg lagi login. atau user yang mau diserah terimakan dokumennya. Kalau 0, backend ambil dari data login
    "share": [1,23,4232,121], //share ke user lain
    "company_id": 23
}
```
Response Body
```
{
    "error": false,
    "message": "Success set document",
    "data": { 
                "document": {
                     "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
                     "name": "Document Job desc Tech",
                     "type": "document",
                     "folder_id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3e", // string kosong atau ada id foldernya
                     "content": {
                         "blocks": [
                             {
                                 "type": "paragraph",
                                 "text": "This is paragraph"
                             }
                         ]
                     }, // block editorjs
                     "timestamp": 1605081795, //second unix timestamp
                     "owner_id": 123, //user yg lagi login. atau user yang mau diserah terimakan dokumennya. Kalau 0, backend ambil dari data login
                     "share": [1,23,4232,121] //share ke user lain
                }
    }
}
```
### Get Detail Document ###
```
GET http://api-gateway.co.id/document-service/document/:document_id
```
Response Body:
```
{
    "error": false,
    "message": "Success get document",
    "data": { 
                "document": {
                     "id": "82b07a6f-60cc-4403-8fd2-329ef0de045d",
                     "name": "Document Job desc Tech",
                     "type": "document",
                     "folder_id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3e", // string kosong atau ada id foldernya
                     "content": {
                         "blocks": [
                             {
                                 "type": "paragraph",
                                 "text": "This is paragraph"
                             }
                         ]
                     }, // block editorjs
                     "timestamp": 1605081795, //second unix timestamp
                     "owner_id": 123, //user yg lagi login. atau user yang mau diserah terimakan dokumennya. Kalau 0, backend ambil dari data login
                     "share": [1,23,4232,121] //share ke user lain
                }
    }
}
```
### Delete Document ###
```
DELETE http://api-gateway.co.id/document-service/document
```
Request body
```
{
        "id": "82b07a6f-60cc-4403-8fd2-329ef0de0d3d" //id document yang dihapus
}
```
Response Body:
```
{
    "error": false,
    "message": "Success delete document",
}
```
