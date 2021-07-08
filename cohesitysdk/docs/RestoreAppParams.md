# RestoreAppParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**OwnerRestoreInfo** | Pointer to [**AppOwnerRestoreInfo**](AppOwnerRestoreInfo.md) |  | [optional] 
**RestoreAppObjectVec** | Pointer to [**[]RestoreAppObject**](RestoreAppObject.md) | The application level objects that needs to be restored. If this vector is populated with exactly one object without its &#39;app_entity&#39;, all the application objects of the owner will be restored. If multiple objects are being restored, the &#39;app_entity&#39; field must be specified for all of them. | [optional] 
**Type** | Pointer to **NullableInt32** | The application environment. | [optional] 

## Methods

### NewRestoreAppParams

`func NewRestoreAppParams() *RestoreAppParams`

NewRestoreAppParams instantiates a new RestoreAppParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAppParamsWithDefaults

`func NewRestoreAppParamsWithDefaults() *RestoreAppParams`

NewRestoreAppParamsWithDefaults instantiates a new RestoreAppParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *RestoreAppParams) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *RestoreAppParams) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *RestoreAppParams) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *RestoreAppParams) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetOwnerRestoreInfo

`func (o *RestoreAppParams) GetOwnerRestoreInfo() AppOwnerRestoreInfo`

GetOwnerRestoreInfo returns the OwnerRestoreInfo field if non-nil, zero value otherwise.

### GetOwnerRestoreInfoOk

`func (o *RestoreAppParams) GetOwnerRestoreInfoOk() (*AppOwnerRestoreInfo, bool)`

GetOwnerRestoreInfoOk returns a tuple with the OwnerRestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRestoreInfo

`func (o *RestoreAppParams) SetOwnerRestoreInfo(v AppOwnerRestoreInfo)`

SetOwnerRestoreInfo sets OwnerRestoreInfo field to given value.

### HasOwnerRestoreInfo

`func (o *RestoreAppParams) HasOwnerRestoreInfo() bool`

HasOwnerRestoreInfo returns a boolean if a field has been set.

### GetRestoreAppObjectVec

`func (o *RestoreAppParams) GetRestoreAppObjectVec() []RestoreAppObject`

GetRestoreAppObjectVec returns the RestoreAppObjectVec field if non-nil, zero value otherwise.

### GetRestoreAppObjectVecOk

`func (o *RestoreAppParams) GetRestoreAppObjectVecOk() (*[]RestoreAppObject, bool)`

GetRestoreAppObjectVecOk returns a tuple with the RestoreAppObjectVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAppObjectVec

`func (o *RestoreAppParams) SetRestoreAppObjectVec(v []RestoreAppObject)`

SetRestoreAppObjectVec sets RestoreAppObjectVec field to given value.

### HasRestoreAppObjectVec

`func (o *RestoreAppParams) HasRestoreAppObjectVec() bool`

HasRestoreAppObjectVec returns a boolean if a field has been set.

### SetRestoreAppObjectVecNil

`func (o *RestoreAppParams) SetRestoreAppObjectVecNil(b bool)`

 SetRestoreAppObjectVecNil sets the value for RestoreAppObjectVec to be an explicit nil

### UnsetRestoreAppObjectVec
`func (o *RestoreAppParams) UnsetRestoreAppObjectVec()`

UnsetRestoreAppObjectVec ensures that no value is present for RestoreAppObjectVec, not even an explicit nil
### GetType

`func (o *RestoreAppParams) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreAppParams) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreAppParams) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreAppParams) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RestoreAppParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RestoreAppParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


