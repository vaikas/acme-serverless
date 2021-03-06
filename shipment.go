package acmeserverless

import "encoding/json"

// ShipmentRequested is the event sent by the Order service when the order is finalized and paid and
// thus ready to be shipped to the customer
type ShipmentRequested struct {
	// Metadata for the event.
	Metadata Metadata `json:"metadata"`

	// Data contains the payload data for the event.
	Data ShipmentRequest `json:"data"`
}

// ShipmentSent is the event sent by the Shipment service when the order is shipped to the customer.
type ShipmentSent struct {
	// Metadata for the event.
	Metadata Metadata `json:"metadata"`

	// Data contains the payload data for the event.
	Data ShipmentData `json:"data"`
}

// ShipmentRequest is the data that the order service emits.
type ShipmentRequest struct {
	// The unique identifier of the order.
	OrderID string `json:"_id"`

	// Delivery is the delivery method that should be used for the shipment.
	Delivery string `json:"delivery"`
}

// ShipmentData is the data the shipment service emits when the shipment is sent.
type ShipmentData struct {
	// The tracking number generated by the shipper.
	TrackingNumber string `json:"trackingNumber"`

	// The unique identifier of the order.
	OrderNumber string `json:"orderNumber"`

	// The current status of the shipment
	Status string `json:"status"`
}

// UnmarshalShipmentRequested parses the JSON-encoded data and stores the result in a
// ShipmentRequestedEvent.
func UnmarshalShipmentRequested(data []byte) (ShipmentRequested, error) {
	var r ShipmentRequested
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal returns the JSON encoding of ShipmentRequested.
func (e *ShipmentRequested) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalShipmentSent parses the JSON-encoded data and stores the result in a
// ShipmentSent.
func UnmarshalShipmentSent(data []byte) (ShipmentSent, error) {
	var r ShipmentSent
	err := json.Unmarshal(data, &r)
	return r, err
}

// Marshal returns the JSON encoding of ShipmentSent.
func (e *ShipmentSent) Marshal() ([]byte, error) {
	return json.Marshal(e)
}
