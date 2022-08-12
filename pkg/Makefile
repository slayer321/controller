client/gen:
	@echo "--> Running code-generator to generate clients"
	# prepare tool code-generator
	@mkdir -p ./tmp/code-generator
	@git clone https://github.com/kubernetes/code-generator.git --branch v0.22.1 --single-branch  ./tmp/code-generator
	# generate client
	./tmp/code-generator/generate-groups.sh "all" github.com/slayer321/crdcreation/pkg/client github.com/slayer321/crdcreation/pkg/apis sachinmaurya.tech:v1alpha1 --go-header-file hack/boilerplate.go.txt
	# check generated client at ./pkg/client
	@cp -r /home/sachinmaurya/go/src/github.com/slayer321/crdcreation/pkg/client/* ./client/
	@rm -rf ./tmp/code-generator



# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) paths="github.com/slayer321/crdcreation/pkg/apis/sachinmaurya.tech/v1alpha1" crd:trivalVersions=true crd:crdVersion=v1 output:crd:artifacts:config=manifests