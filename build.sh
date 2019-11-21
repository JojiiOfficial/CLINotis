#!/bin/bash
go build -o noti && 
if [ -f "/bin/noti" ]; then
	echo \"/bin/noti\" already exists
	echo -n "do you want to continue/overwrite it [y/n] "
	read a
	if [ $a != "y" ]; then
		exit
	fi
fi

sudo mv noti /bin/noti &&
sudo /bin/noti install &&
cat >./checkNewMsg.sh<<EOL
#!/bin/bash
e=\$(/bin/noti check)
if [ ! -z "\$e" ]; then
		echo "You have a new message"
fi
EOL
chmod u+x ./checkNewMsg
