# RecoverAwsParamsRecoverRdsIngestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsTargetParams** | Pointer to [**NullableRecoverRDSPostgresParamsAwsTargetParams**](RecoverRDSPostgresParamsAwsTargetParams.md) |  | [optional] 
**OverwriteDatabase** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**Prefix** | Pointer to **NullableString** | Specifies the prefix to be prepended to the object name after the recovery. | [optional] 
**Suffix** | Pointer to **NullableString** | Specifies the suffix to be appended to the object name after the recovery. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAwsParamsRecoverRdsIngestParams

`func NewRecoverAwsParamsRecoverRdsIngestParams(targetEnvironment string, ) *RecoverAwsParamsRecoverRdsIngestParams`

NewRecoverAwsParamsRecoverRdsIngestParams instantiates a new RecoverAwsParamsRecoverRdsIngestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsParamsRecoverRdsIngestParamsWithDefaults

`func NewRecoverAwsParamsRecoverRdsIngestParamsWithDefaults() *RecoverAwsParamsRecoverRdsIngestParams`

NewRecoverAwsParamsRecoverRdsIngestParamsWithDefaults instantiates a new RecoverAwsParamsRecoverRdsIngestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsTargetParams

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetAwsTargetParams() RecoverRDSPostgresParamsAwsTargetParams`

GetAwsTargetParams returns the AwsTargetParams field if non-nil, zero value otherwise.

### GetAwsTargetParamsOk

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetAwsTargetParamsOk() (*RecoverRDSPostgresParamsAwsTargetParams, bool)`

GetAwsTargetParamsOk returns a tuple with the AwsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetParams

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetAwsTargetParams(v RecoverRDSPostgresParamsAwsTargetParams)`

SetAwsTargetParams sets AwsTargetParams field to given value.

### HasAwsTargetParams

`func (o *RecoverAwsParamsRecoverRdsIngestParams) HasAwsTargetParams() bool`

HasAwsTargetParams returns a boolean if a field has been set.

### SetAwsTargetParamsNil

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetAwsTargetParamsNil(b bool)`

 SetAwsTargetParamsNil sets the value for AwsTargetParams to be an explicit nil

### UnsetAwsTargetParams
`func (o *RecoverAwsParamsRecoverRdsIngestParams) UnsetAwsTargetParams()`

UnsetAwsTargetParams ensures that no value is present for AwsTargetParams, not even an explicit nil
### GetOverwriteDatabase

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetOverwriteDatabase() bool`

GetOverwriteDatabase returns the OverwriteDatabase field if non-nil, zero value otherwise.

### GetOverwriteDatabaseOk

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetOverwriteDatabaseOk() (*bool, bool)`

GetOverwriteDatabaseOk returns a tuple with the OverwriteDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteDatabase

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetOverwriteDatabase(v bool)`

SetOverwriteDatabase sets OverwriteDatabase field to given value.

### HasOverwriteDatabase

`func (o *RecoverAwsParamsRecoverRdsIngestParams) HasOverwriteDatabase() bool`

HasOverwriteDatabase returns a boolean if a field has been set.

### SetOverwriteDatabaseNil

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetOverwriteDatabaseNil(b bool)`

 SetOverwriteDatabaseNil sets the value for OverwriteDatabase to be an explicit nil

### UnsetOverwriteDatabase
`func (o *RecoverAwsParamsRecoverRdsIngestParams) UnsetOverwriteDatabase()`

UnsetOverwriteDatabase ensures that no value is present for OverwriteDatabase, not even an explicit nil
### GetPrefix

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RecoverAwsParamsRecoverRdsIngestParams) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *RecoverAwsParamsRecoverRdsIngestParams) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetSuffix

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *RecoverAwsParamsRecoverRdsIngestParams) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### SetSuffixNil

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetSuffixNil(b bool)`

 SetSuffixNil sets the value for Suffix to be an explicit nil

### UnsetSuffix
`func (o *RecoverAwsParamsRecoverRdsIngestParams) UnsetSuffix()`

UnsetSuffix ensures that no value is present for Suffix, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAwsParamsRecoverRdsIngestParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAwsParamsRecoverRdsIngestParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


