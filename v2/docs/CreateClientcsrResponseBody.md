# CreateClientcsrResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrClient** | Pointer to **NullableString** | Specifies the CSR generated for the client. | [optional] 
**CsrServer** | Pointer to **NullableString** | Specifies the CSR generated for the server. | [optional] 
**FileCsrClient** | Pointer to **NullableString** | Specifies the path to CSR generated for the client | [optional] 
**FileCsrServer** | Pointer to **NullableString** | Specifies the path to CSR generated for the server | [optional] 
**PublicKeyClient** | Pointer to **NullableString** | Specifies the public key generated for this CSR for the client. | [optional] 
**PublicKeyServer** | Pointer to **NullableString** | Specifies the public key generated for this CSR for the server. | [optional] 

## Methods

### NewCreateClientcsrResponseBody

`func NewCreateClientcsrResponseBody() *CreateClientcsrResponseBody`

NewCreateClientcsrResponseBody instantiates a new CreateClientcsrResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientcsrResponseBodyWithDefaults

`func NewCreateClientcsrResponseBodyWithDefaults() *CreateClientcsrResponseBody`

NewCreateClientcsrResponseBodyWithDefaults instantiates a new CreateClientcsrResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrClient

`func (o *CreateClientcsrResponseBody) GetCsrClient() string`

GetCsrClient returns the CsrClient field if non-nil, zero value otherwise.

### GetCsrClientOk

`func (o *CreateClientcsrResponseBody) GetCsrClientOk() (*string, bool)`

GetCsrClientOk returns a tuple with the CsrClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrClient

`func (o *CreateClientcsrResponseBody) SetCsrClient(v string)`

SetCsrClient sets CsrClient field to given value.

### HasCsrClient

`func (o *CreateClientcsrResponseBody) HasCsrClient() bool`

HasCsrClient returns a boolean if a field has been set.

### SetCsrClientNil

`func (o *CreateClientcsrResponseBody) SetCsrClientNil(b bool)`

 SetCsrClientNil sets the value for CsrClient to be an explicit nil

### UnsetCsrClient
`func (o *CreateClientcsrResponseBody) UnsetCsrClient()`

UnsetCsrClient ensures that no value is present for CsrClient, not even an explicit nil
### GetCsrServer

`func (o *CreateClientcsrResponseBody) GetCsrServer() string`

GetCsrServer returns the CsrServer field if non-nil, zero value otherwise.

### GetCsrServerOk

`func (o *CreateClientcsrResponseBody) GetCsrServerOk() (*string, bool)`

GetCsrServerOk returns a tuple with the CsrServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrServer

`func (o *CreateClientcsrResponseBody) SetCsrServer(v string)`

SetCsrServer sets CsrServer field to given value.

### HasCsrServer

`func (o *CreateClientcsrResponseBody) HasCsrServer() bool`

HasCsrServer returns a boolean if a field has been set.

### SetCsrServerNil

`func (o *CreateClientcsrResponseBody) SetCsrServerNil(b bool)`

 SetCsrServerNil sets the value for CsrServer to be an explicit nil

### UnsetCsrServer
`func (o *CreateClientcsrResponseBody) UnsetCsrServer()`

UnsetCsrServer ensures that no value is present for CsrServer, not even an explicit nil
### GetFileCsrClient

`func (o *CreateClientcsrResponseBody) GetFileCsrClient() string`

GetFileCsrClient returns the FileCsrClient field if non-nil, zero value otherwise.

### GetFileCsrClientOk

`func (o *CreateClientcsrResponseBody) GetFileCsrClientOk() (*string, bool)`

GetFileCsrClientOk returns a tuple with the FileCsrClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCsrClient

`func (o *CreateClientcsrResponseBody) SetFileCsrClient(v string)`

SetFileCsrClient sets FileCsrClient field to given value.

### HasFileCsrClient

`func (o *CreateClientcsrResponseBody) HasFileCsrClient() bool`

HasFileCsrClient returns a boolean if a field has been set.

### SetFileCsrClientNil

`func (o *CreateClientcsrResponseBody) SetFileCsrClientNil(b bool)`

 SetFileCsrClientNil sets the value for FileCsrClient to be an explicit nil

### UnsetFileCsrClient
`func (o *CreateClientcsrResponseBody) UnsetFileCsrClient()`

UnsetFileCsrClient ensures that no value is present for FileCsrClient, not even an explicit nil
### GetFileCsrServer

`func (o *CreateClientcsrResponseBody) GetFileCsrServer() string`

GetFileCsrServer returns the FileCsrServer field if non-nil, zero value otherwise.

### GetFileCsrServerOk

`func (o *CreateClientcsrResponseBody) GetFileCsrServerOk() (*string, bool)`

GetFileCsrServerOk returns a tuple with the FileCsrServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCsrServer

`func (o *CreateClientcsrResponseBody) SetFileCsrServer(v string)`

SetFileCsrServer sets FileCsrServer field to given value.

### HasFileCsrServer

`func (o *CreateClientcsrResponseBody) HasFileCsrServer() bool`

HasFileCsrServer returns a boolean if a field has been set.

### SetFileCsrServerNil

`func (o *CreateClientcsrResponseBody) SetFileCsrServerNil(b bool)`

 SetFileCsrServerNil sets the value for FileCsrServer to be an explicit nil

### UnsetFileCsrServer
`func (o *CreateClientcsrResponseBody) UnsetFileCsrServer()`

UnsetFileCsrServer ensures that no value is present for FileCsrServer, not even an explicit nil
### GetPublicKeyClient

`func (o *CreateClientcsrResponseBody) GetPublicKeyClient() string`

GetPublicKeyClient returns the PublicKeyClient field if non-nil, zero value otherwise.

### GetPublicKeyClientOk

`func (o *CreateClientcsrResponseBody) GetPublicKeyClientOk() (*string, bool)`

GetPublicKeyClientOk returns a tuple with the PublicKeyClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyClient

`func (o *CreateClientcsrResponseBody) SetPublicKeyClient(v string)`

SetPublicKeyClient sets PublicKeyClient field to given value.

### HasPublicKeyClient

`func (o *CreateClientcsrResponseBody) HasPublicKeyClient() bool`

HasPublicKeyClient returns a boolean if a field has been set.

### SetPublicKeyClientNil

`func (o *CreateClientcsrResponseBody) SetPublicKeyClientNil(b bool)`

 SetPublicKeyClientNil sets the value for PublicKeyClient to be an explicit nil

### UnsetPublicKeyClient
`func (o *CreateClientcsrResponseBody) UnsetPublicKeyClient()`

UnsetPublicKeyClient ensures that no value is present for PublicKeyClient, not even an explicit nil
### GetPublicKeyServer

`func (o *CreateClientcsrResponseBody) GetPublicKeyServer() string`

GetPublicKeyServer returns the PublicKeyServer field if non-nil, zero value otherwise.

### GetPublicKeyServerOk

`func (o *CreateClientcsrResponseBody) GetPublicKeyServerOk() (*string, bool)`

GetPublicKeyServerOk returns a tuple with the PublicKeyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyServer

`func (o *CreateClientcsrResponseBody) SetPublicKeyServer(v string)`

SetPublicKeyServer sets PublicKeyServer field to given value.

### HasPublicKeyServer

`func (o *CreateClientcsrResponseBody) HasPublicKeyServer() bool`

HasPublicKeyServer returns a boolean if a field has been set.

### SetPublicKeyServerNil

`func (o *CreateClientcsrResponseBody) SetPublicKeyServerNil(b bool)`

 SetPublicKeyServerNil sets the value for PublicKeyServer to be an explicit nil

### UnsetPublicKeyServer
`func (o *CreateClientcsrResponseBody) UnsetPublicKeyServer()`

UnsetPublicKeyServer ensures that no value is present for PublicKeyServer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


