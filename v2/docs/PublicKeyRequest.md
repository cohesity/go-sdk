# PublicKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowType** | **string** | Specifies the workflow initiating the SSH connection. | 

## Methods

### NewPublicKeyRequest

`func NewPublicKeyRequest(workflowType string, ) *PublicKeyRequest`

NewPublicKeyRequest instantiates a new PublicKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeyRequestWithDefaults

`func NewPublicKeyRequestWithDefaults() *PublicKeyRequest`

NewPublicKeyRequestWithDefaults instantiates a new PublicKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowType

`func (o *PublicKeyRequest) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *PublicKeyRequest) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *PublicKeyRequest) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


