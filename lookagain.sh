#! /bin/bash

find . -name '*.sh' | sed 's/.sh//g' | cut -c 3- | rev | cut -d '/' -f1 | rev