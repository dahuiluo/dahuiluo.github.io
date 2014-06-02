#!/bin/sh
rake generate
git add .
git commit -a -m "`date`"
git push origin source
rake deploy
