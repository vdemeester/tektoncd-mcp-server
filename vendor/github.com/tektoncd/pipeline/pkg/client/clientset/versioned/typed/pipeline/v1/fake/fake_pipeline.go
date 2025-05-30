/*
Copyright 2020 The Tekton Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	pipelinev1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakePipelines implements PipelineInterface
type fakePipelines struct {
	*gentype.FakeClientWithList[*v1.Pipeline, *v1.PipelineList]
	Fake *FakeTektonV1
}

func newFakePipelines(fake *FakeTektonV1, namespace string) pipelinev1.PipelineInterface {
	return &fakePipelines{
		gentype.NewFakeClientWithList[*v1.Pipeline, *v1.PipelineList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("pipelines"),
			v1.SchemeGroupVersion.WithKind("Pipeline"),
			func() *v1.Pipeline { return &v1.Pipeline{} },
			func() *v1.PipelineList { return &v1.PipelineList{} },
			func(dst, src *v1.PipelineList) { dst.ListMeta = src.ListMeta },
			func(list *v1.PipelineList) []*v1.Pipeline { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.PipelineList, items []*v1.Pipeline) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
