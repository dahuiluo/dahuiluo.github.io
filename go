#!/bin/sh
rake generate
git add .
git commit -m "`date`"
git push origin source
rake deploy
