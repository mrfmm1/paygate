# Go API client for client

PayGate is a RESTful API enabling first-party Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) transfers to be created without a deep understanding of a full NACHA file specification. First-party transfers initiate at an Originating Depository Financial Institution (ODFI) and are sent off to other Financial Institutions.

Tenants are the largest grouping in PayGate and are typically a vendor who is reselling ACH services or a company making ACH payments themselves. A legal entity is linked off a Tenant as the primary Customer used to KYC and in transfers with the Tenant itself.

An Organization is a grouping within a Tenant which typically represents an entity making ACH transfers. These include clients of an ACH reseller or business accepting payments over ACH. A legal entity is linked off an Organization as the primary Customer used to KYC and in transfers with the Organization itself.

![](https://raw.githubusercontent.com/moov-io/paygate/master/docs/images/tenant-in-paygate.png)


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/moov-io/paygate](https://github.com/moov-io/paygate)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./client"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8082*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MonitorApi* | [**Ping**](docs/MonitorApi.md#ping) | **Get** /ping | Ping PayGate
*OrganizationsApi* | [**CreateOrganization**](docs/OrganizationsApi.md#createorganization) | **Post** /organizations | Create Organization
*OrganizationsApi* | [**GetOrganizations**](docs/OrganizationsApi.md#getorganizations) | **Get** /organizations | Get Organizations
*OrganizationsApi* | [**UpdateOrganization**](docs/OrganizationsApi.md#updateorganization) | **Put** /organizations/{organizationID} | Update Organization
*TenantsApi* | [**GetTenants**](docs/TenantsApi.md#gettenants) | **Get** /tenants | Get Tenants
*TenantsApi* | [**UpdateTenant**](docs/TenantsApi.md#updatetenant) | **Put** /tenants/{tenantID} | Update Tenant
*TransfersApi* | [**AddTransfer**](docs/TransfersApi.md#addtransfer) | **Post** /transfers | Create Transfer
*TransfersApi* | [**DeleteTransferByID**](docs/TransfersApi.md#deletetransferbyid) | **Delete** /transfers/{transferID} | Delete Transfer
*TransfersApi* | [**GetTransferByID**](docs/TransfersApi.md#gettransferbyid) | **Get** /transfers/{transferID} | Get Transfer
*TransfersApi* | [**GetTransfers**](docs/TransfersApi.md#gettransfers) | **Get** /transfers | List Transfers
*ValidationApi* | [**GetAccountMicroDeposits**](docs/ValidationApi.md#getaccountmicrodeposits) | **Get** /accounts/{accountID}/micro-deposits | Get micro-deposits for a specified accountID
*ValidationApi* | [**GetMicroDeposits**](docs/ValidationApi.md#getmicrodeposits) | **Get** /micro-deposits/{microDepositID} | Get micro-deposit information
*ValidationApi* | [**InitiateMicroDeposits**](docs/ValidationApi.md#initiatemicrodeposits) | **Post** /micro-deposits | Create


## Documentation For Models

 - [CreateMicroDeposits](docs/CreateMicroDeposits.md)
 - [CreateOrganization](docs/CreateOrganization.md)
 - [CreateTransfer](docs/CreateTransfer.md)
 - [Destination](docs/Destination.md)
 - [Error](docs/Error.md)
 - [MicroDeposits](docs/MicroDeposits.md)
 - [Organization](docs/Organization.md)
 - [ReturnCode](docs/ReturnCode.md)
 - [Source](docs/Source.md)
 - [Tenant](docs/Tenant.md)
 - [Transfer](docs/Transfer.md)
 - [TransferStatus](docs/TransferStatus.md)
 - [UpdateTenant](docs/UpdateTenant.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



