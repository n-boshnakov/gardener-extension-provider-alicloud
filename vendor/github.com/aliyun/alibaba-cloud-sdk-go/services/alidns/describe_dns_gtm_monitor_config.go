package alidns

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

// DescribeDnsGtmMonitorConfig invokes the alidns.DescribeDnsGtmMonitorConfig API synchronously
func (client *Client) DescribeDnsGtmMonitorConfig(request *DescribeDnsGtmMonitorConfigRequest) (response *DescribeDnsGtmMonitorConfigResponse, err error) {
	response = CreateDescribeDnsGtmMonitorConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDnsGtmMonitorConfigWithChan invokes the alidns.DescribeDnsGtmMonitorConfig API asynchronously
func (client *Client) DescribeDnsGtmMonitorConfigWithChan(request *DescribeDnsGtmMonitorConfigRequest) (<-chan *DescribeDnsGtmMonitorConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsGtmMonitorConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsGtmMonitorConfig(request)
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

// DescribeDnsGtmMonitorConfigWithCallback invokes the alidns.DescribeDnsGtmMonitorConfig API asynchronously
func (client *Client) DescribeDnsGtmMonitorConfigWithCallback(request *DescribeDnsGtmMonitorConfigRequest, callback func(response *DescribeDnsGtmMonitorConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsGtmMonitorConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsGtmMonitorConfig(request)
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

// DescribeDnsGtmMonitorConfigRequest is the request struct for api DescribeDnsGtmMonitorConfig
type DescribeDnsGtmMonitorConfigRequest struct {
	*requests.RpcRequest
	MonitorConfigId string `position:"Query" name:"MonitorConfigId"`
	UserClientIp    string `position:"Query" name:"UserClientIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeDnsGtmMonitorConfigResponse is the response struct for api DescribeDnsGtmMonitorConfig
type DescribeDnsGtmMonitorConfigResponse struct {
	*responses.BaseResponse
	RequestId         string                                    `json:"RequestId" xml:"RequestId"`
	Timeout           int                                       `json:"Timeout" xml:"Timeout"`
	ProtocolType      string                                    `json:"ProtocolType" xml:"ProtocolType"`
	CreateTime        string                                    `json:"CreateTime" xml:"CreateTime"`
	UpdateTime        string                                    `json:"UpdateTime" xml:"UpdateTime"`
	EvaluationCount   int                                       `json:"EvaluationCount" xml:"EvaluationCount"`
	UpdateTimestamp   int64                                     `json:"UpdateTimestamp" xml:"UpdateTimestamp"`
	MonitorExtendInfo string                                    `json:"MonitorExtendInfo" xml:"MonitorExtendInfo"`
	MonitorConfigId   string                                    `json:"MonitorConfigId" xml:"MonitorConfigId"`
	CreateTimestamp   int64                                     `json:"CreateTimestamp" xml:"CreateTimestamp"`
	Interval          int                                       `json:"Interval" xml:"Interval"`
	IspCityNodes      IspCityNodesInDescribeDnsGtmMonitorConfig `json:"IspCityNodes" xml:"IspCityNodes"`
}

// CreateDescribeDnsGtmMonitorConfigRequest creates a request to invoke DescribeDnsGtmMonitorConfig API
func CreateDescribeDnsGtmMonitorConfigRequest() (request *DescribeDnsGtmMonitorConfigRequest) {
	request = &DescribeDnsGtmMonitorConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsGtmMonitorConfig", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDnsGtmMonitorConfigResponse creates a response to parse from DescribeDnsGtmMonitorConfig response
func CreateDescribeDnsGtmMonitorConfigResponse() (response *DescribeDnsGtmMonitorConfigResponse) {
	response = &DescribeDnsGtmMonitorConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
