# CreateReceiver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDepository** | **string** | The depository account to be used by default per transfer. ID must be a valid Receiver Depository account | 
**Email** | **string** | The receivers email address | 
**BirthDate** | [**time.Time**](time.Time.md) | An optional object required for Know Your Customer (KYC) validation of this Receiver. This field is not saved by PayGate.  | [optional] 
**Address** | [**Address**](Address.md) |  | [optional] 
**Metadata** | **string** | Populated into the Entry Detail IndividualName field | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


