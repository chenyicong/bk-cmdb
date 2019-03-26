/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package options

import (
	"configcenter/src/auth/authcenter"
	"configcenter/src/common/core/cc/config"
	"configcenter/src/storage/dal/mongo"
	"configcenter/src/storage/dal/redis"
	"configcenter/src/storage/rpc"

	"github.com/spf13/pflag"
)

//ServerOption define option of server in flags
type ServerOption struct {
	ServConf *config.CCAPIConfig
}

//NewServerOption create a ServerOption object
func NewServerOption() *ServerOption {
	s := ServerOption{
		ServConf: config.NewCCAPIConfig(),
	}

	return &s
}

//AddFlags add flags
func (s *ServerOption) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.ServConf.AddrPort, "addrport", "127.0.0.1:60009", "The ip address and port for the serve on")
	//fs.UintVar(&s.ServConf.Port, "port", 60009, "The port for the serve on")
	fs.StringVar(&s.ServConf.ExConfig, "config", "", "The config path. e.g conf/api.conf")
	fs.StringVar(&s.ServConf.RegDiscover, "regdiscv", "", "hosts of register and discover server. e.g: 127.0.0.1:2181")
}

type Config struct {
	MongoDB mongo.Config
	Redis   redis.Config
	RPC     rpc.ClientConfig
	Auth    authcenter.AuthConfig
}
