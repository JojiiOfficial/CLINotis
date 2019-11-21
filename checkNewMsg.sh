#!/bin/bash
e=$(/bin/noti check)
if [ ! -z "$e" ]; then
		echo "You have a new message"
fi
