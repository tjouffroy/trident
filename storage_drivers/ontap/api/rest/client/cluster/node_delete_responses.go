// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NodeDeleteReader is a Reader for the NodeDelete structure.
type NodeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewNodeDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNodeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodeDeleteOK creates a NodeDeleteOK with default headers values
func NewNodeDeleteOK() *NodeDeleteOK {
	return &NodeDeleteOK{}
}

/*
NodeDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NodeDeleteOK struct {
	Payload *models.NodeJobLinkResponse
}

// IsSuccess returns true when this node delete o k response has a 2xx status code
func (o *NodeDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node delete o k response has a 3xx status code
func (o *NodeDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node delete o k response has a 4xx status code
func (o *NodeDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this node delete o k response has a 5xx status code
func (o *NodeDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this node delete o k response a status code equal to that given
func (o *NodeDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the node delete o k response
func (o *NodeDeleteOK) Code() int {
	return 200
}

func (o *NodeDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/nodes/{uuid}][%d] nodeDeleteOK %s", 200, payload)
}

func (o *NodeDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/nodes/{uuid}][%d] nodeDeleteOK %s", 200, payload)
}

func (o *NodeDeleteOK) GetPayload() *models.NodeJobLinkResponse {
	return o.Payload
}

func (o *NodeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeDeleteAccepted creates a NodeDeleteAccepted with default headers values
func NewNodeDeleteAccepted() *NodeDeleteAccepted {
	return &NodeDeleteAccepted{}
}

/*
NodeDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type NodeDeleteAccepted struct {
	Payload *models.NodeJobLinkResponse
}

// IsSuccess returns true when this node delete accepted response has a 2xx status code
func (o *NodeDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this node delete accepted response has a 3xx status code
func (o *NodeDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this node delete accepted response has a 4xx status code
func (o *NodeDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this node delete accepted response has a 5xx status code
func (o *NodeDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this node delete accepted response a status code equal to that given
func (o *NodeDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the node delete accepted response
func (o *NodeDeleteAccepted) Code() int {
	return 202
}

func (o *NodeDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/nodes/{uuid}][%d] nodeDeleteAccepted %s", 202, payload)
}

func (o *NodeDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/nodes/{uuid}][%d] nodeDeleteAccepted %s", 202, payload)
}

func (o *NodeDeleteAccepted) GetPayload() *models.NodeJobLinkResponse {
	return o.Payload
}

func (o *NodeDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeDeleteDefault creates a NodeDeleteDefault with default headers values
func NewNodeDeleteDefault(code int) *NodeDeleteDefault {
	return &NodeDeleteDefault{
		_statusCode: code,
	}
}

/*
	NodeDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 458755 | Replication service is offline. |
| 458758 | Failed to load job for cluster remove node operation as the job exists. |
| 1179732 | Cannot remove a node in a single-node cluster. |
| 1179735 | Node is not part of a cluster. |
| 1179848 | Operation not supported at the current effective cluster version. |
| 1179849 | Operation failed because the effective cluster version could not be verified. |
| 1182805 | Cannot remove a node from the node network address of the node to be removed. |
| 2293765 | Removing a node only works for nodes not in failover configuration. |
| 2293767 | Node has volumes. Either move or delete them from the node before removing the node. |
| 2293768 | Node is the home node for one or more logical interfaces. |
| 2293769 | Node is the current node for one or more logical interfaces. |
| 2293770 | Node has data logical interfaces configured as target node. |
| 2293787 | Cannot remove node a node with epsilon. |
| 2293788 | Failed to check HA configuration in the cluster. |
| 2293789 | Removing a node only works for nodes not in HA configuration. |
| 2293790 | Failed to check if the node has epsilon. |
| 2293796 | Cluster ring is offline on the node |
| 2293798 | Cannot forcibly remove a node that is online. |
| 2293800 | Node is configured with MetroCluster. |
| 2293801 | Cannot remove node because it has foreign LUN Imports. |
| 2293812 | Node is a member of MetroCluster DR group. |
| 2293813 | Cannot remove a node from the cluster because a controller replacement is in progress. |
| 2293814 | The DELETE operation is not supported until the cluster is upgraded. |
| 2293816 | Cannot remove node because its Storage Encryption devices use authentication keys (AKs) that will not be available to the node after it leaves the cluster. |
| 2293821 | Cannot remove a node that is offline. |
| 10551347 | An update is in progress. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NodeDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this node delete default response has a 2xx status code
func (o *NodeDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this node delete default response has a 3xx status code
func (o *NodeDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this node delete default response has a 4xx status code
func (o *NodeDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this node delete default response has a 5xx status code
func (o *NodeDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this node delete default response a status code equal to that given
func (o *NodeDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the node delete default response
func (o *NodeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NodeDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/nodes/{uuid}][%d] node_delete default %s", o._statusCode, payload)
}

func (o *NodeDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /cluster/nodes/{uuid}][%d] node_delete default %s", o._statusCode, payload)
}

func (o *NodeDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
