# Go API client for quotes

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./quotes"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AssociationsApi* | [**Getcrmv3objectsquotesquoteIdassociationstoObjectTypeGetAll**](docs/AssociationsApi.md#getcrmv3objectsquotesquoteidassociationstoobjecttypegetall) | **Get** /crm/v3/objects/quotes/{quoteId}/associations/{toObjectType} | List associations of a quote by type
*BasicApi* | [**Getcrmv3objectsquotesGetPage**](docs/BasicApi.md#getcrmv3objectsquotesgetpage) | **Get** /crm/v3/objects/quotes | List
*BasicApi* | [**Getcrmv3objectsquotesquoteIdGetById**](docs/BasicApi.md#getcrmv3objectsquotesquoteidgetbyid) | **Get** /crm/v3/objects/quotes/{quoteId} | Read
*BatchApi* | [**Postcrmv3objectsquotesbatchreadRead**](docs/BatchApi.md#postcrmv3objectsquotesbatchreadread) | **Post** /crm/v3/objects/quotes/batch/read | Read a batch of quotes by internal ID, or unique property values
*SearchApi* | [**Postcrmv3objectsquotessearchDoSearch**](docs/SearchApi.md#postcrmv3objectsquotessearchdosearch) | **Post** /crm/v3/objects/quotes/search | 

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

## Author

