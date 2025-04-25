package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/validation/spec"

	api "github.com/tatsuyahayashi/tatope/api/v1"
)

// Generate OpenAPI spec definitions for MPIJob Resource
func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Supply the Tat Operator API version")
	}
	version := os.Args[1]

	filter := func(name string) spec.Ref {
		return spec.MustCreateRef(
			"#/definitions/" + common.EscapeJsonPointer(swaggify(name)))
	}

	oAPIDefs := api.GetOpenAPIDefinitions(filter)
	defs := spec.Definitions{}
	for defName, val := range oAPIDefs {
		//		fmt.Println(val)
		defs[swaggify(defName)] = val.Schema
	}
	swagger := spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Swagger:     "2.0",
			Info:        &spec.Info{InfoProps: spec.InfoProps{Title: "tatoperator", Description: "Python SDK for Tat-Operator", Version: version}},
			Paths:       &spec.Paths{Paths: map[string]spec.PathItem{}},
			Definitions: defs,
			ExternalDocs: &spec.ExternalDocumentation{
				Description: "The tat Operator",
				URL:         "https://tatope.local",
			},
		},
	}

	jsonBytes, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(jsonBytes))
}

// Our strategy here is to replace specific needs with classes we will define
func swaggify(name string) string {

	// These are specific to the Tat Operator
	name = strings.Replace(name, "github.com/tatsuyahayashi/tatope/api/v1.", "", -1)
	// name = strings.Replace(name, "../api/v1alpha2/.", "", -1)
	// name = strings.Replace(name, "./api/v1alpha2/.", "", -1)

	// k8s.io/apimachinery/pkg/apis/meta/v1.Condition -> v1Condition
	//name = strings.Replace(name, "k8s.io/apimachinery/pkg/apis/meta/v1.Condition", "v1Condition", -1)

	// // k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta
	// name = strings.Replace(name, "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta", "v1ListMeta", -1)

	// // k8s.io/apimachinery/pkg/util/intstr.IntOrString -> IntOrString
	// name = strings.Replace(name, "k8s.io/apimachinery/pkg/util/intstr.", "", -1)

	// // k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta
	// name = strings.Replace(name, "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "v1ObjectMeta", -1)

	name = strings.Replace(name, "k8s.io/apimachinery/pkg/apis/meta/", "", -1)
	name = strings.Replace(name, "k8s.io/apimachinery/pkg/runtime/", "", -1)
	name = strings.Replace(name, "k8s.io/apimachinery/pkg/apis/meta/", "", -1)
	name = strings.Replace(name, "k8s.io/apimachinery/pkg/runtime/", "", -1)
	name = strings.Replace(name, "k8s.io/apimachinery/pkg/api/", "", -1)
	name = strings.Replace(name, "k8s.io/kubernetes/pkg/controller/", "", -1)
	name = strings.Replace(name, "k8s.io/client-go/listers/core/", "", -1)
	name = strings.Replace(name, "k8s.io/client-go/util/workqueue", "", -1)
	name = strings.Replace(name, "/", ".", -1)
	return name
}
