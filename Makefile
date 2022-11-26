.DEFAULT: all

.PHONY: all
all: gen

.PHONY: gen
gen:
	hack/update-codegen.sh samplecontroller v1

.PHONY: clean
clean:
	rm -rf ./pkg/client
	rm -f ./pkg/apis/*/*/zz_generated.deepcopy.go