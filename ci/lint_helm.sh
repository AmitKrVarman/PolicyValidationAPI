#!/bin/bash

set -e -u -x

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd | rev | cut -f1 -d/ | rev)

/bin/helm lint ${DIR}/deploy/${DIR}
