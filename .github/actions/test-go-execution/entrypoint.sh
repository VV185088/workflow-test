#!/bin/bash

/get_sum_of_two_values $1 $2;
exitCode=$?;
if [ $exitCode -ne 0 ]; then exit $exitCode; fi
