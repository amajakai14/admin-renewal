#!/bin/bash


#set current directory to where this script is located
cd "$(dirname "$0")"

PGPASSWORD=restaurant psql -d restaurant -h localhost -p 5432 -U restaurant -f ./migrations/corporation_init.sql 



