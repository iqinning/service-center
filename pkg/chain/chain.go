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
package chain

import (
	"fmt"
	"github.com/ServiceComb/service-center/pkg/util"
	"sync"
)

var handlersMap map[string][]Handler

type Chain struct {
	name         string
	handlers     []Handler
	currentIndex int
	mux          sync.Mutex
}

func (c *Chain) Init(chainName string, hs []Handler) {
	c.name = chainName
	c.currentIndex = -1
	if len(hs) > 0 {
		c.handlers = make([]Handler, 0, len(hs))
		copy(c.handlers, hs)
	}
}

func (c *Chain) Name() string {
	return c.name
}

func (c *Chain) doNext(i *Invocation) {
	defer util.RecoverAndReport()

	if c.currentIndex >= len(c.handlers) {
		i.Fail(fmt.Errorf("Over end of chain '%s'", c.name))
		return
	}
	c.currentIndex += 1
	c.handlers[c.currentIndex].Handle(i)
}

func (c *Chain) next(i *Invocation) {
	go c.doNext(i)
}
