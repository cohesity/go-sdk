# ObjectBrowseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Object. | 
**HdfsParams** | Pointer to [**HdfsBrowseRequestParams**](HdfsBrowseRequestParams.md) |  | [optional] 
**PaginationInfo** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewObjectBrowseRequest

`func NewObjectBrowseRequest(environment NullableString, ) *ObjectBrowseRequest`

NewObjectBrowseRequest instantiates a new ObjectBrowseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectBrowseRequestWithDefaults

`func NewObjectBrowseRequestWithDefaults() *ObjectBrowseRequest`

NewObjectBrowseRequestWithDefaults instantiates a new ObjectBrowseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ObjectBrowseRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectBrowseRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectBrowseRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *ObjectBrowseRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectBrowseRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetHdfsParams

`func (o *ObjectBrowseRequest) GetHdfsParams() HdfsBrowseRequestParams`

GetHdfsParams returns the HdfsParams field if non-nil, zero value otherwise.

### GetHdfsParamsOk

`func (o *ObjectBrowseRequest) GetHdfsParamsOk() (*HdfsBrowseRequestParams, bool)`

GetHdfsParamsOk returns a tuple with the HdfsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsParams

`func (o *ObjectBrowseRequest) SetHdfsParams(v HdfsBrowseRequestParams)`

SetHdfsParams sets HdfsParams field to given value.

### HasHdfsParams

`func (o *ObjectBrowseRequest) HasHdfsParams() bool`

HasHdfsParams returns a boolean if a field has been set.

### GetPaginationInfo

`func (o *ObjectBrowseRequest) GetPaginationInfo() PaginationInfo`

GetPaginationInfo returns the PaginationInfo field if non-nil, zero value otherwise.

### GetPaginationInfoOk

`func (o *ObjectBrowseRequest) GetPaginationInfoOk() (*PaginationInfo, bool)`

GetPaginationInfoOk returns a tuple with the PaginationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationInfo

`func (o *ObjectBrowseRequest) SetPaginationInfo(v PaginationInfo)`

SetPaginationInfo sets PaginationInfo field to given value.

### HasPaginationInfo

`func (o *ObjectBrowseRequest) HasPaginationInfo() bool`

HasPaginationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


