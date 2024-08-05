# BucketPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the identifier of the bucket policy. This is a read-only property. | [optional] 
**RawPolicy** | **NullableString** | Specifies the raw JSON string of the store policy. | 
**Version** | Pointer to **NullableString** | Specifies the language syntax rules that are to be used to process the policy. This is a read-only property. | [optional] 

## Methods

### NewBucketPolicy

`func NewBucketPolicy(rawPolicy NullableString, ) *BucketPolicy`

NewBucketPolicy instantiates a new BucketPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketPolicyWithDefaults

`func NewBucketPolicyWithDefaults() *BucketPolicy`

NewBucketPolicyWithDefaults instantiates a new BucketPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BucketPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BucketPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BucketPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BucketPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BucketPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BucketPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRawPolicy

`func (o *BucketPolicy) GetRawPolicy() string`

GetRawPolicy returns the RawPolicy field if non-nil, zero value otherwise.

### GetRawPolicyOk

`func (o *BucketPolicy) GetRawPolicyOk() (*string, bool)`

GetRawPolicyOk returns a tuple with the RawPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPolicy

`func (o *BucketPolicy) SetRawPolicy(v string)`

SetRawPolicy sets RawPolicy field to given value.


### SetRawPolicyNil

`func (o *BucketPolicy) SetRawPolicyNil(b bool)`

 SetRawPolicyNil sets the value for RawPolicy to be an explicit nil

### UnsetRawPolicy
`func (o *BucketPolicy) UnsetRawPolicy()`

UnsetRawPolicy ensures that no value is present for RawPolicy, not even an explicit nil
### GetVersion

`func (o *BucketPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BucketPolicy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BucketPolicy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BucketPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *BucketPolicy) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *BucketPolicy) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


