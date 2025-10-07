#!/bin/sh

geninput(){
	printf 00000000
	printf ff0000ff
	printf 00ff00ff
	printf 0000ffff
}

geninput |
	xxd -r -ps |
	./bytes2rgba2color
