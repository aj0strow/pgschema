output = "bin"
pkg = "github.com/aj0strow/pgschema"
cmd = "pgschema"

targets = [
  %w(linux amd64),
  %w(darwin amd64),
  %w(windows amd64)
]

%x(find bin -maxdepth 1 -type f -delete)

puts "build #{cmd}"

targets.each do |os, arch|
  puts "  build #{os} #{arch}"
  %x(env GOOS=#{os} GOARCH=#{arch} go build #{pkg})
  if os == "windows"
    %x(mv #{cmd}.exe #{output}/#{cmd}-windows-#{arch}.exe)
  else
    %x(tar -cvzf #{cmd}.tar.gz #{cmd})
    %x(rm #{cmd})
    %x(mv #{cmd}.tar.gz #{output}/#{cmd}-#{os}-#{arch}.tar.gz)
  end
end

puts "done"
