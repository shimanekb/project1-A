#!/bin/bash
go build .
mv project1a project1-A
rm -rf storage 2> /dev/null
rm -rf logs.txt 2> /dev/null
rm -rf output.txt 2> /dev/null
