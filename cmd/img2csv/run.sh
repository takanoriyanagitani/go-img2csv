#!/bin/sh

image=~/Downloads/PNG_transparency_demonstration_1.png

output=./sample.d/out.csv

mkdir -p ./sample.d

cat "${image}" |
	./img2csv |
	dd \
		if=/dev/stdin \
		of="${output}" \
		conv=fsync \
		status=progress

file "${image}"
file "${output}"
wc "${output}"
head -1 "${output}" | awk -F, '{ print NF }'
tail -1 "${output}"
