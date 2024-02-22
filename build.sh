#!/bin/bash

make bin

chmod +x bin/kubectl-test

cp bin/kubectl-test /usr/local/bin/
