# TrustedCa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Unique id for the certificate. | [optional] [readonly] 
**Name** | Pointer to **NullableString** | Unique name for the certificate. | [optional] [readonly] 
**IssuedBy** | Pointer to **NullableString** | Specifies the issuer. | [optional] [readonly] 
**IssuedTo** | Pointer to **NullableString** | Specifies whom it was issued to. | [optional] [readonly] 
**IssuedTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp epoch in microseconds when this certificate will start being valid. | [optional] [readonly] 
**ExpirationTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp epoch in microseconds when this certificate will no longer be valid. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | description of the certificate. | [optional] [readonly] 
**RegistrationTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp epoch in microseconds when this certificate was registered on the cluster. | [optional] [readonly] 
**LastValidatedTimeUsecs** | Pointer to **NullableInt64** | Specifies the timestamp epoch in microseconds when this certificate was last validated. | [optional] [readonly] 
**Status** | Pointer to **NullableString** | Validation Status of the certificate. | [optional] [readonly] 

## Methods

### NewTrustedCa

`func NewTrustedCa() *TrustedCa`

NewTrustedCa instantiates a new TrustedCa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedCaWithDefaults

`func NewTrustedCaWithDefaults() *TrustedCa`

NewTrustedCaWithDefaults instantiates a new TrustedCa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustedCa) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustedCa) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustedCa) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrustedCa) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TrustedCa) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TrustedCa) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TrustedCa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustedCa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustedCa) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustedCa) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TrustedCa) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TrustedCa) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIssuedBy

`func (o *TrustedCa) GetIssuedBy() string`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *TrustedCa) GetIssuedByOk() (*string, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *TrustedCa) SetIssuedBy(v string)`

SetIssuedBy sets IssuedBy field to given value.

### HasIssuedBy

`func (o *TrustedCa) HasIssuedBy() bool`

HasIssuedBy returns a boolean if a field has been set.

### SetIssuedByNil

`func (o *TrustedCa) SetIssuedByNil(b bool)`

 SetIssuedByNil sets the value for IssuedBy to be an explicit nil

### UnsetIssuedBy
`func (o *TrustedCa) UnsetIssuedBy()`

UnsetIssuedBy ensures that no value is present for IssuedBy, not even an explicit nil
### GetIssuedTo

`func (o *TrustedCa) GetIssuedTo() string`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *TrustedCa) GetIssuedToOk() (*string, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *TrustedCa) SetIssuedTo(v string)`

SetIssuedTo sets IssuedTo field to given value.

### HasIssuedTo

`func (o *TrustedCa) HasIssuedTo() bool`

HasIssuedTo returns a boolean if a field has been set.

### SetIssuedToNil

`func (o *TrustedCa) SetIssuedToNil(b bool)`

 SetIssuedToNil sets the value for IssuedTo to be an explicit nil

### UnsetIssuedTo
`func (o *TrustedCa) UnsetIssuedTo()`

UnsetIssuedTo ensures that no value is present for IssuedTo, not even an explicit nil
### GetIssuedTimeUsecs

`func (o *TrustedCa) GetIssuedTimeUsecs() int64`

GetIssuedTimeUsecs returns the IssuedTimeUsecs field if non-nil, zero value otherwise.

### GetIssuedTimeUsecsOk

`func (o *TrustedCa) GetIssuedTimeUsecsOk() (*int64, bool)`

GetIssuedTimeUsecsOk returns a tuple with the IssuedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTimeUsecs

`func (o *TrustedCa) SetIssuedTimeUsecs(v int64)`

SetIssuedTimeUsecs sets IssuedTimeUsecs field to given value.

### HasIssuedTimeUsecs

`func (o *TrustedCa) HasIssuedTimeUsecs() bool`

HasIssuedTimeUsecs returns a boolean if a field has been set.

### SetIssuedTimeUsecsNil

`func (o *TrustedCa) SetIssuedTimeUsecsNil(b bool)`

 SetIssuedTimeUsecsNil sets the value for IssuedTimeUsecs to be an explicit nil

### UnsetIssuedTimeUsecs
`func (o *TrustedCa) UnsetIssuedTimeUsecs()`

UnsetIssuedTimeUsecs ensures that no value is present for IssuedTimeUsecs, not even an explicit nil
### GetExpirationTimeUsecs

`func (o *TrustedCa) GetExpirationTimeUsecs() int64`

GetExpirationTimeUsecs returns the ExpirationTimeUsecs field if non-nil, zero value otherwise.

### GetExpirationTimeUsecsOk

`func (o *TrustedCa) GetExpirationTimeUsecsOk() (*int64, bool)`

GetExpirationTimeUsecsOk returns a tuple with the ExpirationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimeUsecs

`func (o *TrustedCa) SetExpirationTimeUsecs(v int64)`

SetExpirationTimeUsecs sets ExpirationTimeUsecs field to given value.

### HasExpirationTimeUsecs

`func (o *TrustedCa) HasExpirationTimeUsecs() bool`

HasExpirationTimeUsecs returns a boolean if a field has been set.

### SetExpirationTimeUsecsNil

`func (o *TrustedCa) SetExpirationTimeUsecsNil(b bool)`

 SetExpirationTimeUsecsNil sets the value for ExpirationTimeUsecs to be an explicit nil

### UnsetExpirationTimeUsecs
`func (o *TrustedCa) UnsetExpirationTimeUsecs()`

UnsetExpirationTimeUsecs ensures that no value is present for ExpirationTimeUsecs, not even an explicit nil
### GetDescription

`func (o *TrustedCa) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrustedCa) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrustedCa) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrustedCa) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TrustedCa) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TrustedCa) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRegistrationTimeUsecs

`func (o *TrustedCa) GetRegistrationTimeUsecs() int64`

GetRegistrationTimeUsecs returns the RegistrationTimeUsecs field if non-nil, zero value otherwise.

### GetRegistrationTimeUsecsOk

`func (o *TrustedCa) GetRegistrationTimeUsecsOk() (*int64, bool)`

GetRegistrationTimeUsecsOk returns a tuple with the RegistrationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTimeUsecs

`func (o *TrustedCa) SetRegistrationTimeUsecs(v int64)`

SetRegistrationTimeUsecs sets RegistrationTimeUsecs field to given value.

### HasRegistrationTimeUsecs

`func (o *TrustedCa) HasRegistrationTimeUsecs() bool`

HasRegistrationTimeUsecs returns a boolean if a field has been set.

### SetRegistrationTimeUsecsNil

`func (o *TrustedCa) SetRegistrationTimeUsecsNil(b bool)`

 SetRegistrationTimeUsecsNil sets the value for RegistrationTimeUsecs to be an explicit nil

### UnsetRegistrationTimeUsecs
`func (o *TrustedCa) UnsetRegistrationTimeUsecs()`

UnsetRegistrationTimeUsecs ensures that no value is present for RegistrationTimeUsecs, not even an explicit nil
### GetLastValidatedTimeUsecs

`func (o *TrustedCa) GetLastValidatedTimeUsecs() int64`

GetLastValidatedTimeUsecs returns the LastValidatedTimeUsecs field if non-nil, zero value otherwise.

### GetLastValidatedTimeUsecsOk

`func (o *TrustedCa) GetLastValidatedTimeUsecsOk() (*int64, bool)`

GetLastValidatedTimeUsecsOk returns a tuple with the LastValidatedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidatedTimeUsecs

`func (o *TrustedCa) SetLastValidatedTimeUsecs(v int64)`

SetLastValidatedTimeUsecs sets LastValidatedTimeUsecs field to given value.

### HasLastValidatedTimeUsecs

`func (o *TrustedCa) HasLastValidatedTimeUsecs() bool`

HasLastValidatedTimeUsecs returns a boolean if a field has been set.

### SetLastValidatedTimeUsecsNil

`func (o *TrustedCa) SetLastValidatedTimeUsecsNil(b bool)`

 SetLastValidatedTimeUsecsNil sets the value for LastValidatedTimeUsecs to be an explicit nil

### UnsetLastValidatedTimeUsecs
`func (o *TrustedCa) UnsetLastValidatedTimeUsecs()`

UnsetLastValidatedTimeUsecs ensures that no value is present for LastValidatedTimeUsecs, not even an explicit nil
### GetStatus

`func (o *TrustedCa) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrustedCa) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrustedCa) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrustedCa) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TrustedCa) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TrustedCa) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


