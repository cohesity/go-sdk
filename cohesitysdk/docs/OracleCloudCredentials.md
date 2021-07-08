# OracleCloudCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **NullableString** | Specifies access key to connect to Oracle S3 Compatible vault account. | [optional] 
**Region** | Pointer to **NullableString** | Specifies the region for Oracle S3 Compatible vault account. | [optional] 
**SecretAccessKey** | Pointer to **NullableString** | Specifies the secret access key for Oracle S3 Compatible vault account. | [optional] 
**Tenant** | Pointer to **NullableString** | Specifies the tenant which is part of the REST endpoints for Oracle S3 compatible vaults. | [optional] 
**TierType** | Pointer to **NullableString** | Specifies the storage class of Oracle vault. OracleTierType specifies the storage class for Oracle. &#39;kOracleTierStandard&#39; indicates a tier type of Oracle properties that requires fast, immediate and frequent access. &#39;kOracleTierArchive&#39; indicates a tier type of Oracle properties that is rarely accesed and preserved for long times. | [optional] 

## Methods

### NewOracleCloudCredentials

`func NewOracleCloudCredentials() *OracleCloudCredentials`

NewOracleCloudCredentials instantiates a new OracleCloudCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleCloudCredentialsWithDefaults

`func NewOracleCloudCredentialsWithDefaults() *OracleCloudCredentials`

NewOracleCloudCredentialsWithDefaults instantiates a new OracleCloudCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *OracleCloudCredentials) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *OracleCloudCredentials) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *OracleCloudCredentials) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *OracleCloudCredentials) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyIdNil

`func (o *OracleCloudCredentials) SetAccessKeyIdNil(b bool)`

 SetAccessKeyIdNil sets the value for AccessKeyId to be an explicit nil

### UnsetAccessKeyId
`func (o *OracleCloudCredentials) UnsetAccessKeyId()`

UnsetAccessKeyId ensures that no value is present for AccessKeyId, not even an explicit nil
### GetRegion

`func (o *OracleCloudCredentials) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OracleCloudCredentials) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OracleCloudCredentials) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OracleCloudCredentials) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *OracleCloudCredentials) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *OracleCloudCredentials) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSecretAccessKey

`func (o *OracleCloudCredentials) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *OracleCloudCredentials) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *OracleCloudCredentials) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *OracleCloudCredentials) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### SetSecretAccessKeyNil

`func (o *OracleCloudCredentials) SetSecretAccessKeyNil(b bool)`

 SetSecretAccessKeyNil sets the value for SecretAccessKey to be an explicit nil

### UnsetSecretAccessKey
`func (o *OracleCloudCredentials) UnsetSecretAccessKey()`

UnsetSecretAccessKey ensures that no value is present for SecretAccessKey, not even an explicit nil
### GetTenant

`func (o *OracleCloudCredentials) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *OracleCloudCredentials) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *OracleCloudCredentials) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *OracleCloudCredentials) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *OracleCloudCredentials) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *OracleCloudCredentials) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTierType

`func (o *OracleCloudCredentials) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *OracleCloudCredentials) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *OracleCloudCredentials) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *OracleCloudCredentials) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### SetTierTypeNil

`func (o *OracleCloudCredentials) SetTierTypeNil(b bool)`

 SetTierTypeNil sets the value for TierType to be an explicit nil

### UnsetTierType
`func (o *OracleCloudCredentials) UnsetTierType()`

UnsetTierType ensures that no value is present for TierType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


