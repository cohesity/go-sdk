# ClusterPackageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMsg** | Pointer to **string** | Error message if package is not available.  | [optional] 
**Status** | **string** | Status of the package * &#x60;Available&#x60; - Package is available for use. * &#x60;DownloadFailed&#x60; - Package download has failed.  | 

## Methods

### NewClusterPackageStatus

`func NewClusterPackageStatus(status string, ) *ClusterPackageStatus`

NewClusterPackageStatus instantiates a new ClusterPackageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPackageStatusWithDefaults

`func NewClusterPackageStatusWithDefaults() *ClusterPackageStatus`

NewClusterPackageStatusWithDefaults instantiates a new ClusterPackageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMsg

`func (o *ClusterPackageStatus) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *ClusterPackageStatus) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *ClusterPackageStatus) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *ClusterPackageStatus) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterPackageStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterPackageStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterPackageStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


