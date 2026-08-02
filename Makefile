all: run
run:
	list ./src/ > ./src/SUMMARY.md 
	mdbook build
	./rss > ./docs/feed
	git add . && git commit -m "update"
.PHONY: all
