/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterTemplateRevisionsGetter has a method to return a ClusterTemplateRevisionInterface.
// A group's client should implement this interface.
type ClusterTemplateRevisionsGetter interface {
	ClusterTemplateRevisions(namespace string) ClusterTemplateRevisionInterface
}

// ClusterTemplateRevisionInterface has methods to work with ClusterTemplateRevision resources.
type ClusterTemplateRevisionInterface interface {
	Create(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.CreateOptions) (*v3.ClusterTemplateRevision, error)
	Update(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.UpdateOptions) (*v3.ClusterTemplateRevision, error)
	UpdateStatus(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.UpdateOptions) (*v3.ClusterTemplateRevision, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.ClusterTemplateRevision, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ClusterTemplateRevisionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterTemplateRevision, err error)
	ClusterTemplateRevisionExpansion
}

// clusterTemplateRevisions implements ClusterTemplateRevisionInterface
type clusterTemplateRevisions struct {
	client rest.Interface
	ns     string
}

// newClusterTemplateRevisions returns a ClusterTemplateRevisions
func newClusterTemplateRevisions(c *ManagementV3Client, namespace string) *clusterTemplateRevisions {
	return &clusterTemplateRevisions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterTemplateRevision, and returns the corresponding clusterTemplateRevision object, and an error if there is any.
func (c *clusterTemplateRevisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterTemplateRevision, err error) {
	result = &v3.ClusterTemplateRevision{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterTemplateRevisions that match those selectors.
func (c *clusterTemplateRevisions) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterTemplateRevisionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.ClusterTemplateRevisionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterTemplateRevisions.
func (c *clusterTemplateRevisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterTemplateRevision and creates it.  Returns the server's representation of the clusterTemplateRevision, and an error, if there is any.
func (c *clusterTemplateRevisions) Create(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.CreateOptions) (result *v3.ClusterTemplateRevision, err error) {
	result = &v3.ClusterTemplateRevision{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTemplateRevision).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterTemplateRevision and updates it. Returns the server's representation of the clusterTemplateRevision, and an error, if there is any.
func (c *clusterTemplateRevisions) Update(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.UpdateOptions) (result *v3.ClusterTemplateRevision, err error) {
	result = &v3.ClusterTemplateRevision{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		Name(clusterTemplateRevision.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTemplateRevision).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterTemplateRevisions) UpdateStatus(ctx context.Context, clusterTemplateRevision *v3.ClusterTemplateRevision, opts v1.UpdateOptions) (result *v3.ClusterTemplateRevision, err error) {
	result = &v3.ClusterTemplateRevision{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		Name(clusterTemplateRevision.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTemplateRevision).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterTemplateRevision and deletes it. Returns an error if one occurs.
func (c *clusterTemplateRevisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterTemplateRevisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterTemplateRevision.
func (c *clusterTemplateRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterTemplateRevision, err error) {
	result = &v3.ClusterTemplateRevision{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clustertemplaterevisions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
