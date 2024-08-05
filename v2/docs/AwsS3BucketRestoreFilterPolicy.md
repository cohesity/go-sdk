# AwsS3BucketRestoreFilterPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeList** | Pointer to **[]string** | List of include prefixes that need to be recovered. | [optional] 

## Methods

### NewAwsS3BucketRestoreFilterPolicy

`func NewAwsS3BucketRestoreFilterPolicy() *AwsS3BucketRestoreFilterPolicy`

NewAwsS3BucketRestoreFilterPolicy instantiates a new AwsS3BucketRestoreFilterPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsS3BucketRestoreFilterPolicyWithDefaults

`func NewAwsS3BucketRestoreFilterPolicyWithDefaults() *AwsS3BucketRestoreFilterPolicy`

NewAwsS3BucketRestoreFilterPolicyWithDefaults instantiates a new AwsS3BucketRestoreFilterPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeList

`func (o *AwsS3BucketRestoreFilterPolicy) GetIncludeList() []string`

GetIncludeList returns the IncludeList field if non-nil, zero value otherwise.

### GetIncludeListOk

`func (o *AwsS3BucketRestoreFilterPolicy) GetIncludeListOk() (*[]string, bool)`

GetIncludeListOk returns a tuple with the IncludeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeList

`func (o *AwsS3BucketRestoreFilterPolicy) SetIncludeList(v []string)`

SetIncludeList sets IncludeList field to given value.

### HasIncludeList

`func (o *AwsS3BucketRestoreFilterPolicy) HasIncludeList() bool`

HasIncludeList returns a boolean if a field has been set.

### SetIncludeListNil

`func (o *AwsS3BucketRestoreFilterPolicy) SetIncludeListNil(b bool)`

 SetIncludeListNil sets the value for IncludeList to be an explicit nil

### UnsetIncludeList
`func (o *AwsS3BucketRestoreFilterPolicy) UnsetIncludeList()`

UnsetIncludeList ensures that no value is present for IncludeList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


