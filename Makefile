all: run
run:
	tree=$(tree -tf --noreport -d 'src' -I '.*|img|list|*.xml' --charset ascii $1 | sed -e 's/| \+/  /g' -e 's/[|`]-\+/ */g' -e 's:\(* \)\(\(.*/\)\([^/]\+\)\):\1[\4](\2):g') printf "# Project files \n\n${tree}" > ./src/SUMMARY.md
	mdbook build
	./rss > ./docs/feed
	git add . && git commit -m "update"
.PHONY: all
