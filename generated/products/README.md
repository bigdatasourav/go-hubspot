# Go API client for products

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./products"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AssociationsApi* | [**Deletecrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeArchive**](docs/AssociationsApi.md#deletecrmv3objectsproductsproductidassociationstoobjecttypetoobjectidassociationtypearchive) | **Delete** /crm/v3/objects/products/{productId}/associations/{toObjectType}/{toObjectId}/{associationType} | Remove an association between two products
*AssociationsApi* | [**Getcrmv3objectsproductsproductIdassociationstoObjectTypeGetAll**](docs/AssociationsApi.md#getcrmv3objectsproductsproductidassociationstoobjecttypegetall) | **Get** /crm/v3/objects/products/{productId}/associations/{toObjectType} | List associations of a product by type
*AssociationsApi* | [**Putcrmv3objectsproductsproductIdassociationstoObjectTypetoObjectIdassociationTypeCreate**](docs/AssociationsApi.md#putcrmv3objectsproductsproductidassociationstoobjecttypetoobjectidassociationtypecreate) | **Put** /crm/v3/objects/products/{productId}/associations/{toObjectType}/{toObjectId}/{associationType} | Associate a product with another object
*BasicApi* | [**Deletecrmv3objectsproductsproductIdArchive**](docs/BasicApi.md#deletecrmv3objectsproductsproductidarchive) | **Delete** /crm/v3/objects/products/{productId} | Archive
*BasicApi* | [**Getcrmv3objectsproductsGetPage**](docs/BasicApi.md#getcrmv3objectsproductsgetpage) | **Get** /crm/v3/objects/products | List
*BasicApi* | [**Getcrmv3objectsproductsproductIdGetById**](docs/BasicApi.md#getcrmv3objectsproductsproductidgetbyid) | **Get** /crm/v3/objects/products/{productId} | Read
*BasicApi* | [**Patchcrmv3objectsproductsproductIdUpdate**](docs/BasicApi.md#patchcrmv3objectsproductsproductidupdate) | **Patch** /crm/v3/objects/products/{productId} | Update
*BasicApi* | [**Postcrmv3objectsproductsCreate**](docs/BasicApi.md#postcrmv3objectsproductscreate) | **Post** /crm/v3/objects/products | Create
*BatchApi* | [**Postcrmv3objectsproductsbatcharchiveArchive**](docs/BatchApi.md#postcrmv3objectsproductsbatcharchivearchive) | **Post** /crm/v3/objects/products/batch/archive | Archive a batch of products by ID
*BatchApi* | [**Postcrmv3objectsproductsbatchcreateCreate**](docs/BatchApi.md#postcrmv3objectsproductsbatchcreatecreate) | **Post** /crm/v3/objects/products/batch/create | Create a batch of products
*BatchApi* | [**Postcrmv3objectsproductsbatchreadRead**](docs/BatchApi.md#postcrmv3objectsproductsbatchreadread) | **Post** /crm/v3/objects/products/batch/read | Read a batch of products by internal ID, or unique property values
*BatchApi* | [**Postcrmv3objectsproductsbatchupdateUpdate**](docs/BatchApi.md#postcrmv3objectsproductsbatchupdateupdate) | **Post** /crm/v3/objects/products/batch/update | Update a batch of products
*SearchApi* | [**Postcrmv3objectsproductssearchDoSearch**](docs/SearchApi.md#postcrmv3objectsproductssearchdosearch) | **Post** /crm/v3/objects/products/search | 

## Documentation For Models

 - [AssociatedId](docs/AssociatedId.md)
 - [BatchInputSimplePublicObjectBatchInput](docs/BatchInputSimplePublicObjectBatchInput.md)
 - [BatchInputSimplePublicObjectId](docs/BatchInputSimplePublicObjectId.md)
 - [BatchInputSimplePublicObjectInput](docs/BatchInputSimplePublicObjectInput.md)
 - [BatchReadInputSimplePublicObjectId](docs/BatchReadInputSimplePublicObjectId.md)
 - [BatchResponseSimplePublicObject](docs/BatchResponseSimplePublicObject.md)
 - [BatchResponseSimplePublicObjectWithErrors](docs/BatchResponseSimplePublicObjectWithErrors.md)
 - [CollectionResponseAssociatedId](docs/CollectionResponseAssociatedId.md)
 - [CollectionResponseAssociatedIdForwardPaging](docs/CollectionResponseAssociatedIdForwardPaging.md)
 - [CollectionResponseSimplePublicObjectWithAssociationsForwardPaging](docs/CollectionResponseSimplePublicObjectWithAssociationsForwardPaging.md)
 - [CollectionResponseWithTotalSimplePublicObjectForwardPaging](docs/CollectionResponseWithTotalSimplePublicObjectForwardPaging.md)
 - [ErrorCategory](docs/ErrorCategory.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [Filter](docs/Filter.md)
 - [FilterGroup](docs/FilterGroup.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [ModelError](docs/ModelError.md)
 - [NextPage](docs/NextPage.md)
 - [Paging](docs/Paging.md)
 - [PreviousPage](docs/PreviousPage.md)
 - [PublicObjectSearchRequest](docs/PublicObjectSearchRequest.md)
 - [SimplePublicObject](docs/SimplePublicObject.md)
 - [SimplePublicObjectBatchInput](docs/SimplePublicObjectBatchInput.md)
 - [SimplePublicObjectId](docs/SimplePublicObjectId.md)
 - [SimplePublicObjectInput](docs/SimplePublicObjectInput.md)
 - [SimplePublicObjectWithAssociations](docs/SimplePublicObjectWithAssociations.md)
 - [StandardError](docs/StandardError.md)

## Documentation For Authorization

## hapikey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## oauth2_legacy
- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **e-commerce**: e-commerce

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

## Author

