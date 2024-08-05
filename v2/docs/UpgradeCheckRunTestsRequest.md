# UpgradeCheckRunTestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | Pointer to **string** | Type of upgrade checks(pre/post) to run | [optional] 

## Methods

### NewUpgradeCheckRunTestsRequest

`func NewUpgradeCheckRunTestsRequest() *UpgradeCheckRunTestsRequest`

NewUpgradeCheckRunTestsRequest instantiates a new UpgradeCheckRunTestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeCheckRunTestsRequestWithDefaults

`func NewUpgradeCheckRunTestsRequestWithDefaults() *UpgradeCheckRunTestsRequest`

NewUpgradeCheckRunTestsRequestWithDefaults instantiates a new UpgradeCheckRunTestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestType

`func (o *UpgradeCheckRunTestsRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *UpgradeCheckRunTestsRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *UpgradeCheckRunTestsRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *UpgradeCheckRunTestsRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


