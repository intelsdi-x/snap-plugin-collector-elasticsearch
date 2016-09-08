default:
	$(MAKE) deps
	$(MAKE) all
deps:
	bash -c "./scripts/deps.sh"
test:
	bash -c "./scripts/test.sh $(TEST_TYPE)"
test-unit:
	bash -c "./scripts/test.sh unit"
test-integration:
	bash -c "./scripts/test.sh integration"
check:
	$(MAKE) test
all:
	bash -c "./scripts/build.sh $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))"