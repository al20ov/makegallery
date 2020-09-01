SRC		=		main.go

OUT		=		makegallery

all:
	go build -o $(OUT) $(SRC)

clean:
	rm -f $(OUT) index.html

.PHONY: all clean
