# CassandraSourceRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JmxCredentials** | Pointer to [**NullableCassandraSourceRegistrationParamsAllOfJmxCredentials**](CassandraSourceRegistrationParamsAllOfJmxCredentials.md) |  | [optional] 
**CassandraCredentials** | Pointer to [**NullableCassandraSourceRegistrationParamsAllOfCassandraCredentials**](CassandraSourceRegistrationParamsAllOfCassandraCredentials.md) |  | [optional] 
**DataCenterNames** | Pointer to **[]string** | Data centers for this cluster. | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Principal for the kerberos connection. (This is required only if your Cassandra has Kerberos authentication. Please refer to the user guide.) | [optional] 
**DseSolrInfo** | Pointer to [**DSESolrInfo**](DSESolrInfo.md) |  | [optional] 

## Methods

### NewCassandraSourceRegistrationParamsAllOf

`func NewCassandraSourceRegistrationParamsAllOf() *CassandraSourceRegistrationParamsAllOf`

NewCassandraSourceRegistrationParamsAllOf instantiates a new CassandraSourceRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSourceRegistrationParamsAllOfWithDefaults

`func NewCassandraSourceRegistrationParamsAllOfWithDefaults() *CassandraSourceRegistrationParamsAllOf`

NewCassandraSourceRegistrationParamsAllOfWithDefaults instantiates a new CassandraSourceRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJmxCredentials

`func (o *CassandraSourceRegistrationParamsAllOf) GetJmxCredentials() CassandraSourceRegistrationParamsAllOfJmxCredentials`

GetJmxCredentials returns the JmxCredentials field if non-nil, zero value otherwise.

### GetJmxCredentialsOk

`func (o *CassandraSourceRegistrationParamsAllOf) GetJmxCredentialsOk() (*CassandraSourceRegistrationParamsAllOfJmxCredentials, bool)`

GetJmxCredentialsOk returns a tuple with the JmxCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxCredentials

`func (o *CassandraSourceRegistrationParamsAllOf) SetJmxCredentials(v CassandraSourceRegistrationParamsAllOfJmxCredentials)`

SetJmxCredentials sets JmxCredentials field to given value.

### HasJmxCredentials

`func (o *CassandraSourceRegistrationParamsAllOf) HasJmxCredentials() bool`

HasJmxCredentials returns a boolean if a field has been set.

### SetJmxCredentialsNil

`func (o *CassandraSourceRegistrationParamsAllOf) SetJmxCredentialsNil(b bool)`

 SetJmxCredentialsNil sets the value for JmxCredentials to be an explicit nil

### UnsetJmxCredentials
`func (o *CassandraSourceRegistrationParamsAllOf) UnsetJmxCredentials()`

UnsetJmxCredentials ensures that no value is present for JmxCredentials, not even an explicit nil
### GetCassandraCredentials

`func (o *CassandraSourceRegistrationParamsAllOf) GetCassandraCredentials() CassandraSourceRegistrationParamsAllOfCassandraCredentials`

GetCassandraCredentials returns the CassandraCredentials field if non-nil, zero value otherwise.

### GetCassandraCredentialsOk

`func (o *CassandraSourceRegistrationParamsAllOf) GetCassandraCredentialsOk() (*CassandraSourceRegistrationParamsAllOfCassandraCredentials, bool)`

GetCassandraCredentialsOk returns a tuple with the CassandraCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraCredentials

`func (o *CassandraSourceRegistrationParamsAllOf) SetCassandraCredentials(v CassandraSourceRegistrationParamsAllOfCassandraCredentials)`

SetCassandraCredentials sets CassandraCredentials field to given value.

### HasCassandraCredentials

`func (o *CassandraSourceRegistrationParamsAllOf) HasCassandraCredentials() bool`

HasCassandraCredentials returns a boolean if a field has been set.

### SetCassandraCredentialsNil

`func (o *CassandraSourceRegistrationParamsAllOf) SetCassandraCredentialsNil(b bool)`

 SetCassandraCredentialsNil sets the value for CassandraCredentials to be an explicit nil

### UnsetCassandraCredentials
`func (o *CassandraSourceRegistrationParamsAllOf) UnsetCassandraCredentials()`

UnsetCassandraCredentials ensures that no value is present for CassandraCredentials, not even an explicit nil
### GetDataCenterNames

`func (o *CassandraSourceRegistrationParamsAllOf) GetDataCenterNames() []string`

GetDataCenterNames returns the DataCenterNames field if non-nil, zero value otherwise.

### GetDataCenterNamesOk

`func (o *CassandraSourceRegistrationParamsAllOf) GetDataCenterNamesOk() (*[]string, bool)`

GetDataCenterNamesOk returns a tuple with the DataCenterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenterNames

`func (o *CassandraSourceRegistrationParamsAllOf) SetDataCenterNames(v []string)`

SetDataCenterNames sets DataCenterNames field to given value.

### HasDataCenterNames

`func (o *CassandraSourceRegistrationParamsAllOf) HasDataCenterNames() bool`

HasDataCenterNames returns a boolean if a field has been set.

### GetKerberosPrincipal

`func (o *CassandraSourceRegistrationParamsAllOf) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *CassandraSourceRegistrationParamsAllOf) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *CassandraSourceRegistrationParamsAllOf) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *CassandraSourceRegistrationParamsAllOf) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *CassandraSourceRegistrationParamsAllOf) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *CassandraSourceRegistrationParamsAllOf) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetDseSolrInfo

`func (o *CassandraSourceRegistrationParamsAllOf) GetDseSolrInfo() DSESolrInfo`

GetDseSolrInfo returns the DseSolrInfo field if non-nil, zero value otherwise.

### GetDseSolrInfoOk

`func (o *CassandraSourceRegistrationParamsAllOf) GetDseSolrInfoOk() (*DSESolrInfo, bool)`

GetDseSolrInfoOk returns a tuple with the DseSolrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseSolrInfo

`func (o *CassandraSourceRegistrationParamsAllOf) SetDseSolrInfo(v DSESolrInfo)`

SetDseSolrInfo sets DseSolrInfo field to given value.

### HasDseSolrInfo

`func (o *CassandraSourceRegistrationParamsAllOf) HasDseSolrInfo() bool`

HasDseSolrInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


