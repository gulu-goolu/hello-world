#!/bin/bash
# run_tests

DIR=.
if [ $# -eq 1 ]; then
	DIR=$1
fi

# test d$ir not exists
if [ ! -d $DIR ]; then
	exit 1
fi

function run_test {
	cd $1
	
	local count=0
	local files=$(ls .)
	for f in $files; do
		if [ -d $f ]; then
			run_test $f
			# check result
			if [ ! $? -eq 0 ]; then
				return $?
			fi
		elif [ $f == *test.go ]; then
			count=$[ $count + 1 ]
		fi
	done

	# do test
	if [ $count -gt 0 ]; then
		echo $(pwd)
		go test
	fi

	# back to parent directory
	local r=$?
	cd ..
	
	return $r
}

run_test $DIR
