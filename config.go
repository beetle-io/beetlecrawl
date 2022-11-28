// Copyright 2022 beetlecrawl Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package beetlecrawl

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type (
	AppConfig struct {
		Downloader DownloaderConfig `yaml:"downloader"`
	}

	DownloaderConfig struct {
		MaxRetry int `yaml:"max_retry"`
	}
)

func LoadConfig(filePath string) (*AppConfig, error) {
	confFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var appConf AppConfig
	if err = yaml.Unmarshal(confFile, &appConf); err != nil {
		return nil, err
	}
	return &appConf, nil
}
