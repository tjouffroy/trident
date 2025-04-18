// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/scheme"
)

// TridentBackendsGetter has a method to return a TridentBackendInterface.
// A group's client should implement this interface.
type TridentBackendsGetter interface {
	TridentBackends(namespace string) TridentBackendInterface
}

// TridentBackendInterface has methods to work with TridentBackend resources.
type TridentBackendInterface interface {
	Create(ctx context.Context, tridentBackend *v1.TridentBackend, opts metav1.CreateOptions) (*v1.TridentBackend, error)
	Update(ctx context.Context, tridentBackend *v1.TridentBackend, opts metav1.UpdateOptions) (*v1.TridentBackend, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentBackend, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentBackendList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentBackend, err error)
	TridentBackendExpansion
}

// tridentBackends implements TridentBackendInterface
type tridentBackends struct {
	client rest.Interface
	ns     string
}

// newTridentBackends returns a TridentBackends
func newTridentBackends(c *TridentV1Client, namespace string) *tridentBackends {
	return &tridentBackends{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentBackend, and returns the corresponding tridentBackend object, and an error if there is any.
func (c *tridentBackends) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentBackend, err error) {
	result = &v1.TridentBackend{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentbackends").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentBackends that match those selectors.
func (c *tridentBackends) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentBackendList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentBackendList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentbackends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentBackends.
func (c *tridentBackends) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentbackends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentBackend and creates it.  Returns the server's representation of the tridentBackend, and an error, if there is any.
func (c *tridentBackends) Create(ctx context.Context, tridentBackend *v1.TridentBackend, opts metav1.CreateOptions) (result *v1.TridentBackend, err error) {
	result = &v1.TridentBackend{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentbackends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentBackend).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentBackend and updates it. Returns the server's representation of the tridentBackend, and an error, if there is any.
func (c *tridentBackends) Update(ctx context.Context, tridentBackend *v1.TridentBackend, opts metav1.UpdateOptions) (result *v1.TridentBackend, err error) {
	result = &v1.TridentBackend{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentbackends").
		Name(tridentBackend.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentBackend).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentBackend and deletes it. Returns an error if one occurs.
func (c *tridentBackends) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentbackends").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentBackends) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentbackends").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentBackend.
func (c *tridentBackends) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentBackend, err error) {
	result = &v1.TridentBackend{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentbackends").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
