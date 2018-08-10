/*
Copyright 2018 Compose, Zalando SE

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/zalando-incubator/postgres-operator/pkg/apis/acid.zalan.do/v1"
	scheme "github.com/zalando-incubator/postgres-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperatorConfigurationsGetter has a method to return a OperatorConfigurationInterface.
// A group's client should implement this interface.
type OperatorConfigurationsGetter interface {
	OperatorConfigurations(namespace string) OperatorConfigurationInterface
}

// OperatorConfigurationInterface has methods to work with OperatorConfiguration resources.
type OperatorConfigurationInterface interface {
	Create(*v1.OperatorConfiguration) (*v1.OperatorConfiguration, error)
	Update(*v1.OperatorConfiguration) (*v1.OperatorConfiguration, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.OperatorConfiguration, error)
	List(opts metav1.ListOptions) (*v1.OperatorConfigurationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.OperatorConfiguration, err error)
	OperatorConfigurationExpansion
}

// operatorConfigurations implements OperatorConfigurationInterface
type operatorConfigurations struct {
	client rest.Interface
	ns     string
}

// newOperatorConfigurations returns a OperatorConfigurations
func newOperatorConfigurations(c *AcidV1Client, namespace string) *operatorConfigurations {
	return &operatorConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the operatorConfiguration, and returns the corresponding operatorConfiguration object, and an error if there is any.
func (c *operatorConfigurations) Get(name string, options metav1.GetOptions) (result *v1.OperatorConfiguration, err error) {
	result = &v1.OperatorConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OperatorConfigurations that match those selectors.
func (c *operatorConfigurations) List(opts metav1.ListOptions) (result *v1.OperatorConfigurationList, err error) {
	result = &v1.OperatorConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operatorConfigurations.
func (c *operatorConfigurations) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a operatorConfiguration and creates it.  Returns the server's representation of the operatorConfiguration, and an error, if there is any.
func (c *operatorConfigurations) Create(operatorConfiguration *v1.OperatorConfiguration) (result *v1.OperatorConfiguration, err error) {
	result = &v1.OperatorConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		Body(operatorConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a operatorConfiguration and updates it. Returns the server's representation of the operatorConfiguration, and an error, if there is any.
func (c *operatorConfigurations) Update(operatorConfiguration *v1.OperatorConfiguration) (result *v1.OperatorConfiguration, err error) {
	result = &v1.OperatorConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		Name(operatorConfiguration.Name).
		Body(operatorConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the operatorConfiguration and deletes it. Returns an error if one occurs.
func (c *operatorConfigurations) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operatorConfigurations) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched operatorConfiguration.
func (c *operatorConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.OperatorConfiguration, err error) {
	result = &v1.OperatorConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("operatorconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
