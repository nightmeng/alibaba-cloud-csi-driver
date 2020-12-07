package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeCapacityReservations invokes the ecs.DescribeCapacityReservations API synchronously
func (client *Client) DescribeCapacityReservations(request *DescribeCapacityReservationsRequest) (response *DescribeCapacityReservationsResponse, err error) {
	response = CreateDescribeCapacityReservationsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCapacityReservationsWithChan invokes the ecs.DescribeCapacityReservations API asynchronously
func (client *Client) DescribeCapacityReservationsWithChan(request *DescribeCapacityReservationsRequest) (<-chan *DescribeCapacityReservationsResponse, <-chan error) {
	responseChan := make(chan *DescribeCapacityReservationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCapacityReservations(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeCapacityReservationsWithCallback invokes the ecs.DescribeCapacityReservations API asynchronously
func (client *Client) DescribeCapacityReservationsWithCallback(request *DescribeCapacityReservationsRequest, callback func(response *DescribeCapacityReservationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCapacityReservationsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCapacityReservations(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeCapacityReservationsRequest is the request struct for api DescribeCapacityReservations
type DescribeCapacityReservationsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Platform              string           `position:"Query" name:"Platform"`
	NextToken             string           `position:"Query" name:"NextToken"`
	InstanceType          string           `position:"Query" name:"InstanceType"`
	InstanceChargeType    string           `position:"Query" name:"InstanceChargeType"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	PrivatePoolOptionsIds string           `position:"Query" name:"PrivatePoolOptions.Ids"`
	MaxResults            requests.Integer `position:"Query" name:"MaxResults"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	PackageType           string           `position:"Query" name:"PackageType"`
	Status                string           `position:"Query" name:"Status"`
}

// DescribeCapacityReservationsResponse is the response struct for api DescribeCapacityReservations
type DescribeCapacityReservationsResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	NextToken              string                 `json:"NextToken" xml:"NextToken"`
	MaxResults             int                    `json:"MaxResults" xml:"MaxResults"`
	TotalCount             int                    `json:"TotalCount" xml:"TotalCount"`
	CapacityReservationSet CapacityReservationSet `json:"CapacityReservationSet" xml:"CapacityReservationSet"`
}

// CreateDescribeCapacityReservationsRequest creates a request to invoke DescribeCapacityReservations API
func CreateDescribeCapacityReservationsRequest() (request *DescribeCapacityReservationsRequest) {
	request = &DescribeCapacityReservationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeCapacityReservations", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCapacityReservationsResponse creates a response to parse from DescribeCapacityReservations response
func CreateDescribeCapacityReservationsResponse() (response *DescribeCapacityReservationsResponse) {
	response = &DescribeCapacityReservationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
