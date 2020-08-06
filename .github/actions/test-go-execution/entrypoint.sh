#!/bin/bash

/push-to-service-now $1 "$2";
exitCode=$?;
if [ $exitCode -ne 0 ]; then exit $exitCode; fi
