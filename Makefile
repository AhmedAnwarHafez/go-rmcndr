start:
	go run main.go & \
	tailwindcss -i ./public/input.css -o ./public/output.css --watch
