#### Code Generation using code-generator repo

```
execDir=/home/sachinmaurya/go/pkg/mod/k8s.io/code-generator@v0.22.1
```


```
"${execDir}"/generate-groups.sh "all" github.com/slayer321/crdcreation/pkg/client github.com/slayer321/crdcreation/pkg/apis sachinmaurya.tech:v1alpha1 --go-header-file  "${execDir}"/hack/boilerplate.go.txt 
```


#### CRD generator command using controller-gen

```
controller-gen paths=github.com/slayer321/crdcreation/pkg/apis/sachinmaurya.tech/v1alpha1 crd:trivalVersions=true crd:crdVersion=v1 output:crd:artifacts:config=manifests

```