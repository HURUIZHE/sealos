/*
Copyright 2022 cuisongliu@qq.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package binary

import (
	"fmt"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/labring/sealos/pkg/utils/exec"
	"github.com/labring/sealos/pkg/utils/logger"

	"github.com/labring/sealos/pkg/image/types"
)

type RegistryService struct {
}

func (*RegistryService) Login(domain, username, passwd string) error {
	return exec.Cmd("bash", "-c", fmt.Sprintf("buildah login --tls-verify=false --username %s --password %s %s", username, passwd, domain))
}
func (*RegistryService) Logout(domain string) error {
	return exec.Cmd("bash", "-c", fmt.Sprintf("buildah logout %s", domain))
}
func (*RegistryService) Pull(platform v1.Platform, images ...string) error {
	platformCmd := fmt.Sprintf(" --platform %s", fmt.Sprintf("%s/%s", platform.OS, platform.Architecture))
	logger.Info("pull images %v for platform is %s", images, fmt.Sprintf("%s/%s", platform.OS, platform.Architecture))

	for _, image := range images {
		if err := exec.Cmd("bash", "-c", fmt.Sprintf("buildah pull --tls-verify=false %s %s", platformCmd, image)); err != nil {
			return err
		}
	}
	return nil
}

func (*RegistryService) Push(image string) error {
	return exec.Cmd("bash", "-c", fmt.Sprintf("buildah push --tls-verify=false %s", image))
}

func NewRegistryService() (types.RegistryService, error) {
	return &RegistryService{}, nil
}
