#!/bin/bash

set -e -u -x

find policyValidationAPI/ -name "Dockerfile*" \
	-exec echo "Starting linting for policy-validation-api" {} \; \
	-exec dockerfile_lint -p -f {} \;

echo "Linted all files with success"
