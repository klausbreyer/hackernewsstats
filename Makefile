
start:
	go run .


# Tunnel to the raspi to access the mysql db.
tunnel:
	 ssh -N -L 3336:127.0.0.1:3306 pi@10.0.0.3
