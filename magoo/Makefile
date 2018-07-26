cmd = magoo
dirs = service store
test = tests

$(dirs): 
	cd $@ && make 

build: $(dirs) $(cmd)
	cd $< && make build

test: $(dirs) $(cmd)
	cd $< && make test

magoo: magoo/$(cmd)
	cd $(cmd) && make $@

run: $(cmd) 
	cd $(cmd) && make run

install: $(src)
	cd $< && go install 

dbuild:
	docker build -t magoo .

drun:
	docker run -it --rm -v /srv/magfs:/srv/magfs --expose 1199 -p 1199:1199 --name magoo magoo

deploy:
	gcloud app deploy dispatch.yml

.PHONY: $(dirs)
