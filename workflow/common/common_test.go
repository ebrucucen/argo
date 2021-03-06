package common

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/stretchr/testify/assert"
)

func TestUnstructuredHasCompletedLabel(t *testing.T) {
	noLabel := &unstructured.Unstructured{}
	assert.False(t, UnstructuredHasCompletedLabel(noLabel))

	label := &unstructured.Unstructured{Object: map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				LabelKeyCompleted: "true",
			},
		},
	}}
	assert.True(t, UnstructuredHasCompletedLabel(label))

	falseLabel := &unstructured.Unstructured{Object: map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				LabelKeyCompleted: "false",
			},
		},
	}}
	assert.False(t, UnstructuredHasCompletedLabel(falseLabel))

	unknownObject := "hello"
	assert.Panics(t, func() { UnstructuredHasCompletedLabel(unknownObject) })
}
