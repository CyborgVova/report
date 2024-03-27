#!/bin/bash

docker build -t report .
docker run --rm -d --name report
