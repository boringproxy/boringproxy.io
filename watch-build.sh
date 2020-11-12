#!/bin/bash

ls pages/* partials/* | entr ./ssg.js $1
