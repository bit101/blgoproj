default:
	@rm -rf frames
	@mkdir frames
	@go run main.go

clean:
	@rm -rf frames
	@rm -f out.png out.gif out.mp4
