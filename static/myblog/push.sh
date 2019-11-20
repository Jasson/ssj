#!/usr/bin/env bash
make clean && make
git add .
git commit -a -m "update"
git push origin master
