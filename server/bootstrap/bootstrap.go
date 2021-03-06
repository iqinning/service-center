//Copyright 2017 Huawei Technologies Co., Ltd
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
package bootstrap

import _ "github.com/ServiceComb/service-center/server/core" // initialize
// cipher
import _ "github.com/ServiceComb/service-center/server/plugin/infra/security/plain"

// registry
import _ "github.com/ServiceComb/service-center/server/core/registry/etcd"
import _ "github.com/ServiceComb/service-center/server/core/registry/embededetcd"

// v3
import _ "github.com/ServiceComb/service-center/server/rest/controller/v3"

// quota
import _ "github.com/ServiceComb/service-center/server/plugin/infra/quota/buildin"
import _ "github.com/ServiceComb/service-center/server/plugin/infra/quota/unlimit"

import (
	"github.com/ServiceComb/service-center/pkg/util"
	"github.com/ServiceComb/service-center/server/interceptor"
	"github.com/ServiceComb/service-center/server/interceptor/access"
	"github.com/ServiceComb/service-center/server/interceptor/cors"
	"github.com/ServiceComb/service-center/server/interceptor/maxbody"
	"github.com/ServiceComb/service-center/server/interceptor/ratelimiter"
)

func init() {
	util.Logger().Info("BootStrap Huawei Enterprise Edition")

	interceptor.InterceptFunc(interceptor.ACCESS_PHASE, ratelimiter.Intercept)
	interceptor.InterceptFunc(interceptor.ACCESS_PHASE, access.Intercept)
	interceptor.InterceptFunc(interceptor.ACCESS_PHASE, cors.Intercept)

	interceptor.InterceptFunc(interceptor.CONTENT_PHASE, maxbody.Intercept)

	interceptor.InterceptFunc(interceptor.LOG_PHASE, access.Log)
}
