all: run
run:
	mdbook build
	./rss > ./docs/feed
	git add . && git commit -m "update"
.PHONY: all
