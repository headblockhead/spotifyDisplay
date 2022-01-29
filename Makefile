build-pi:
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./CLOCK

deploy: build-pi
	-ssh pi@spotipy.local sudo killall CLOCK
	scp ./CLOCK pi@spotipy.local:/home/pi/CLOCK
	ssh pi@spotipy.local sudo ./CLOCK
