#!/bin/sh

# If number of arguments is not (numerically) equal to 1
if [ "$#" -ne 1 ]
then
  echo "> Please only provide 1 argument with the name of the entity to generate."
  exit 1
fi

# Create entity
go run entgo.io/ent/cmd/ent init "$1"
