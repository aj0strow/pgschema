package=github.com/aj0strow/pgschema
targets=(linux/amd64)

for i in ${targets[@]}; do
	cmd=$(basename ${package})
	os=$(dirname ${i})
	arch=$(basename ${i})
	fname="${cmd}-${os}-${arch}"
	rm "${fname}" 2>/dev/null || true
	rm "${fname}.tar.gz" 2>/dev/null || true
	GOOS=${os} GOARCH=${arch} go build -o "${fname}" "${package}"
	tar -cz -f "${fname}.tar.gz" "${fname}"
	rm "${fname}"
done
