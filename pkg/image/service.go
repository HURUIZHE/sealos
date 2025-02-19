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

package image

import (
	"github.com/labring/sealos/pkg/image/binary"
	buildah_cluster "github.com/labring/sealos/pkg/image/buildah/cluster"
	buildah_image "github.com/labring/sealos/pkg/image/buildah/image"
	"github.com/labring/sealos/pkg/image/buildah/registry"
	"github.com/labring/sealos/pkg/image/types"
	"github.com/labring/sealos/pkg/utils/logger"
)

func NewClusterService() (types.ClusterService, error) {
	if ok, err := initBuildah(); err == nil && ok {
		logger.Debug("setting binary build to cluster service")
		return binary.NewClusterService()
	}
	logger.Debug("setting sdk build to cluster service")
	return buildah_cluster.NewClusterService()
}

func NewRegistryService() (types.RegistryService, error) {
	if ok, err := initBuildah(); err == nil && ok {
		logger.Debug("setting binary build to registry service")
		return binary.NewRegistryService()
	}
	logger.Debug("setting sdk build to registry service")
	return registry.NewRegistryService()
}

func NewImageService() (types.ImageService, error) {
	if ok, err := initBuildah(); err == nil && ok {
		logger.Debug("setting binary build to image service")
		return binary.NewImageService()
	}
	logger.Debug("setting sdk build to image service")
	return buildah_image.NewImageService()
}
