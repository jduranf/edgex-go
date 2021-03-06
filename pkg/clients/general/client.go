/*******************************************************************************
 * Copyright 2019 Dell Inc.
 * Copyright 2019 Joan Duran
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package general

import (
	"context"

	"github.com/jduranf/edgex-go/pkg/clients"
)

type GeneralClient interface {
	FetchConfiguration(ctx context.Context) (string, error)
	FetchMetrics(ctx context.Context) (string, error)
}

type generalRestClient struct {
	url string
}

func NewGeneralClient(url string) GeneralClient {
	gc := generalRestClient{url: url}
	return &gc
}

// FetchConfiguration fetch configuration information from the service.
func (gc *generalRestClient) FetchConfiguration(ctx context.Context) (string, error) {
	body, err := clients.GetRequest(gc.url+clients.ApiConfigRoute, ctx)
	return string(body), err
}

// FetchMetrics fetch metrics information from the service.
func (gc *generalRestClient) FetchMetrics(ctx context.Context) (string, error) {
	body, err := clients.GetRequest(gc.url+clients.ApiMetricsRoute, ctx)
	return string(body), err
}
