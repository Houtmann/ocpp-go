package ocpp16

import (
	"reflect"
)

// -------------------- Remote Stop Transaction (CS -> CP) --------------------

// The field definition of the RemoteStopTransaction request payload sent by the Central System to the Charge Point.
type RemoteStopTransactionRequest struct {
	TransactionId int `json:"transactionId" validate:"gte=0"`
}

// This field definition of the RemoteStopTransaction confirmation payload, sent by the Charge Point to the Central System in response to a RemoteStopTransactionRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type RemoteStopTransactionConfirmation struct {
	Status RemoteStartStopStatus `json:"status" validate:"required,remoteStartStopStatus"`
}

// Central System can request a Charge Point to stop a transaction by sending a RemoteStopTransactionRequest to Charge Point with the identifier of the transaction.
// Charge Point SHALL reply with RemoteStopTransactionConfirmation and a status indicating whether it has accepted the request and a transaction with the given transactionId is ongoing and will be stopped.
// This remote request to stop a transaction is equal to a local action to stop a transaction.
// Therefore, the transaction SHALL be stopped, The Charge Point SHALL send a StopTransactionRequest and, if applicable, unlock the connector.
type RemoteStopTransactionFeature struct{}

func (f RemoteStopTransactionFeature) GetFeatureName() string {
	return RemoteStopTransactionFeatureName
}

func (f RemoteStopTransactionFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(RemoteStopTransactionRequest{})
}

func (f RemoteStopTransactionFeature) GetConfirmationType() reflect.Type {
	return reflect.TypeOf(RemoteStopTransactionConfirmation{})
}

func (r RemoteStopTransactionRequest) GetFeatureName() string {
	return RemoteStopTransactionFeatureName
}

func (c RemoteStopTransactionConfirmation) GetFeatureName() string {
	return RemoteStopTransactionFeatureName
}

// Creates a new RemoteStopTransactionRequest, containing all required fields. There are no optional fields for this message.
func NewRemoteStopTransactionRequest(transactionId int) *RemoteStopTransactionRequest {
	return &RemoteStopTransactionRequest{TransactionId: transactionId}
}

// Creates a new RemoteStopTransactionConfirmation, containing all required fields. There are no optional fields for this message.
func NewRemoteStopTransactionConfirmation(status RemoteStartStopStatus) *RemoteStopTransactionConfirmation {
	return &RemoteStopTransactionConfirmation{Status: status}
}
