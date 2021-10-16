module kmodules.xyz/jsonpatch-gen

go 1.12

require (
	github.com/evanphx/json-patch v4.11.0+incompatible // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	gomodules.xyz/jsonpatch/v3 v3.0.1-0.20200322101522-b31198dbf14f
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/satori/go.uuid => github.com/gofrs/uuid v4.0.0+incompatible

replace github.com/dgrijalva/jwt-go => github.com/gomodules/jwt v3.2.2+incompatible

replace github.com/golang-jwt/jwt => github.com/golang-jwt/jwt v3.2.2+incompatible

replace github.com/form3tech-oss/jwt-go => github.com/form3tech-oss/jwt-go v3.2.5+incompatible

replace helm.sh/helm/v3 => github.com/kubepack/helm/v3 v3.6.1-0.20210518225915-c3e0ce48dd1b

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.21.2-0.20210716212718-83e5493ac170
