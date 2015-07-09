VERSION=0.1.2

buildmake:
	rm -rf RepoTsar
	mkdir -p RepoTsar/bin/
	rm -rf pkg1
	mkdir pkg1
	cp ../../../../RepoTsar ./RepoTsar/bin/
	pkgbuild --root RepoTsar \
    		--identifier org.RepoTsar.command \
    		--version $(VERSION) \
    		--ownership recommended \
    		pkg1/RepoTsar.pkg
	productbuild --distribution ./distrobution.xml \
              --resources ./resources \
              --package-path ./pkg1 \
              --version $(VERSION) \
              ../RepoTsar-v$(VERSION)-amd64-osx10.8.pkg
