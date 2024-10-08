Cohesity Go SDK
===============

## Overview
The *Cohesity Go SDK*  provides an easy-to-use language binding to
harness the power of *Cohesity REST APIs* in your Go applications.

## Table of contents :scroll:
 - [Getting Started](#get-started)
 - [Cluster Version and Compatibility Matrix](#compatibility-matrix)
 - [How to use](#howto)
 - [Suggestions and Feedback](#suggest)
 

## <a name="get-started"></a> Let's get started :hammer_and_pick:

### Installation
```
go get github.com/cohesity/go-sdk
```
## <a name="compatibility-matrix"></a> Compatibility Matrix

|Cluster Version| SDK Version|
|---|--|
|7.1.2_u2|7.1.2201|
## <a name="howto"></a> How to Use: :mag_right:
This SDK exposes all the functionality provided by *Cohesity REST API*.

v1:
```go
import(
  apiClient "github.com/cohesity/go-sdk/v1/client"
  "github.com/cohesity/go-sdk/v1/client/access_tokens"
  "github.com/cohesity/go-sdk/v1/models"
  "github.com/cohesity/go-sdk/v1/client/protection_sources"
  "github.com/go-openapi/strfmt"

)
client := apiClient.NewHTTPClientWithConfig(strfmt.Default, apiClient.DefaultTransportConfig().WithHost("cluster_ip"))
username := "user"
password := "password"
domain := "LOCAL"
body := &models.AccessTokenCredential{
  Username: &username,
  Password: &password,
  Domain:   &domain,
}
resp, err := client.AccessTokens.GenerateAccessToken(access_tokens.NewGenerateAccessTokenParams().WithBody(body), nil)
bearerTokenAuth := httptransport.BearerToken(*resp.Payload.AccessToken) # the bearertoken is the authentication we use to access the APIs
sourceResp, err := client.ProtectionSources.ListProtectionSources(protection_sources.NewListProtectionSourcesParams(), bearerTokenAuth) #example api
```
v2:
```go
import(
  apiClient "github.com/cohesity/go-sdk/v2/client"
  "github.com/cohesity/go-sdk/v2/client/access_token"
  "github.com/cohesity/go-sdk/v2/models"
  "github.com/cohesity/go-sdk/v2/client/protection_group"
  "github.com/go-openapi/strfmt"

)
client := apiClient.NewHTTPClientWithConfig(strfmt.Default, apiClient.DefaultTransportConfig().WithHost("cluster_ip"))
username := "user"
password := "password"
domain := "LOCAL"
body := &models.CreateAccessTokenRequestParams{
  Username: &username,
  Password: &password,
  Domain:   &domain,
}
resp, err := client.AccessToken.CreateAccessToken(access_token.NewCreateAccessTokenParams().WithBody(body), nil)
bearerTokenAuth := httptransport.BearerToken(*resp.Payload.AccessToken) # the bearertoken is the authentication we use to access the APIs
protectionGroupResp, err := client.ProtectionGroup.GetProtectionGroups(protection_group.NewGetProtectionGroupsParams(), bearerTokenAuth) #example api
```

## <a name ="suggest"></a> Questions or Feedback :raised_hand:

We would love to hear from you. Please send your questions and feedback via Cohesity support portal: *https://www.cohesity.com/support/*
