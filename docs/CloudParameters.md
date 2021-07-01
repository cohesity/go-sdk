# CloudParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverToCloud** | Pointer to **NullableBool** | Specifies whether the Protection Sources in this Protection Job will be failed over to Cloud. This flag is applicable to backup on-prem Sources. | [optional] 

## Methods

### NewCloudParameters

`func NewCloudParameters() *CloudParameters`

NewCloudParameters instantiates a new CloudParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudParametersWithDefaults

`func NewCloudParametersWithDefaults() *CloudParameters`

NewCloudParametersWithDefaults instantiates a new CloudParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailoverToCloud

`func (o *CloudParameters) GetFailoverToCloud() bool`

GetFailoverToCloud returns the FailoverToCloud field if non-nil, zero value otherwise.

### GetFailoverToCloudOk

`func (o *CloudParameters) GetFailoverToCloudOk() (*bool, bool)`

GetFailoverToCloudOk returns a tuple with the FailoverToCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverToCloud

`func (o *CloudParameters) SetFailoverToCloud(v bool)`

SetFailoverToCloud sets FailoverToCloud field to given value.

### HasFailoverToCloud

`func (o *CloudParameters) HasFailoverToCloud() bool`

HasFailoverToCloud returns a boolean if a field has been set.

### SetFailoverToCloudNil

`func (o *CloudParameters) SetFailoverToCloudNil(b bool)`

 SetFailoverToCloudNil sets the value for FailoverToCloud to be an explicit nil

### UnsetFailoverToCloud
`func (o *CloudParameters) UnsetFailoverToCloud()`

UnsetFailoverToCloud ensures that no value is present for FailoverToCloud, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


