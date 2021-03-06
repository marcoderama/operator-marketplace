package operatorsource_test

import (
	"fmt"

	marketplace "github.com/operator-framework/operator-marketplace/pkg/apis/operators/v1"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func helperGetContextLogger() *log.Entry {
	return log.NewEntry(log.New())
}

func helperNewOperatorSourceWithEndpoint(namespace, name, endpointType string) *marketplace.OperatorSource {
	return &marketplace.OperatorSource{
		TypeMeta: metav1.TypeMeta{
			APIVersion: fmt.Sprintf("%s/%s",
				marketplace.SchemeGroupVersion.Group, marketplace.SchemeGroupVersion.Version),
			Kind: marketplace.OperatorSourceKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},

		Spec: marketplace.OperatorSourceSpec{
			Type:     endpointType,
			Endpoint: "http://localhost:5000/cnr",
		},
	}
}

func helperNewOperatorSourceWithPhase(namespace, name, phase string) *marketplace.OperatorSource {
	return &marketplace.OperatorSource{
		TypeMeta: metav1.TypeMeta{
			APIVersion: fmt.Sprintf("%s/%s",
				marketplace.SchemeGroupVersion.Group, marketplace.SchemeGroupVersion.Version),
			Kind: marketplace.OperatorSourceKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},

		Spec: marketplace.OperatorSourceSpec{
			Type:     "appregistry",
			Endpoint: "http://localhost:5000/cnr",
		},

		Status: marketplace.OperatorSourceStatus{
			CurrentPhase: marketplace.ObjectPhase{
				Phase: marketplace.Phase{
					Name: phase,
				},
			},
		},
	}
}

func helperNewCatalogSourceConfig(namespace, name string) *marketplace.CatalogSourceConfig {
	return &marketplace.CatalogSourceConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: fmt.Sprintf("%s/%s",
				marketplace.SchemeGroupVersion.Group, marketplace.SchemeGroupVersion.Version),
			Kind: marketplace.CatalogSourceConfigKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func helperNewCatalogSourceConfigWithLabels(namespace, name string, opsrcLabels map[string]string) *marketplace.CatalogSourceConfig {
	csc := helperNewCatalogSourceConfig(namespace, name)

	// This is the default label that should get added to CatalogSourceConfig.
	labels := map[string]string{
		"opsrc-datastore": "true",
	}

	for key, value := range opsrcLabels {
		labels[key] = value
	}

	csc.SetLabels(labels)

	return csc
}
