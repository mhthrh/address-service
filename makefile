set-app-path:
	sudo chmod +x ./script/set.sh
	zsh ./script/set.sh
	#sh ./script/set.sh
update-common-lib:
	sudo chmod +x ./script/update-lib.sh
	zsh ./script/update-lib.sh
	#sh ./script/update-lib.sh
test_app:
	go test ./test/... -v -bench . -failfast -cover -count=1