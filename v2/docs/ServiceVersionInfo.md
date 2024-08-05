# ServiceVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCommitTime** | Pointer to **string** | LastCommitTime of the service. | [optional] 
**ServiceName** | Pointer to **string** | Name of the service. | [optional] 
**ServiceVersion** | Pointer to **string** | Version of the service. | [optional] 

## Methods

### NewServiceVersionInfo

`func NewServiceVersionInfo() *ServiceVersionInfo`

NewServiceVersionInfo instantiates a new ServiceVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVersionInfoWithDefaults

`func NewServiceVersionInfoWithDefaults() *ServiceVersionInfo`

NewServiceVersionInfoWithDefaults instantiates a new ServiceVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCommitTime

`func (o *ServiceVersionInfo) GetLastCommitTime() string`

GetLastCommitTime returns the LastCommitTime field if non-nil, zero value otherwise.

### GetLastCommitTimeOk

`func (o *ServiceVersionInfo) GetLastCommitTimeOk() (*string, bool)`

GetLastCommitTimeOk returns a tuple with the LastCommitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommitTime

`func (o *ServiceVersionInfo) SetLastCommitTime(v string)`

SetLastCommitTime sets LastCommitTime field to given value.

### HasLastCommitTime

`func (o *ServiceVersionInfo) HasLastCommitTime() bool`

HasLastCommitTime returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceVersionInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceVersionInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceVersionInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceVersionInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ServiceVersionInfo) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ServiceVersionInfo) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ServiceVersionInfo) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ServiceVersionInfo) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


