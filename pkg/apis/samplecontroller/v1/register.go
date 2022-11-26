/*
@Time : 2022/11/26 14:47
@Author : lianyz
@Description :
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	api "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   "samplecontroller",
	Version: "v1",
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&At{},
		&AtList{})

	api.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil

}
