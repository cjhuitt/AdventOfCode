total = 0
path = []
files = Hash.new(0)

def command(path, cmd)
  bits = cmd.split()
  return path if bits[0] != "cd"
  case bits[1]
  when "/"
    []
  when ".."
    path.pop
    path
  else
    path.push bits[1]
  end
  path
end

def full_path(path)
  '/' + path.join('/')
end

def parse_contents(path, line, files)
  bits = line.split()
  case bits[0]
  when "dir"
    # nothing
  else
    p = path.clone
    until p.empty?
      files[full_path(p)] += bits[0].to_i
      p.pop
    end
    files["/"] += bits[0].to_i
  end
end

File.foreach("input.txt") do |line|
  case line[0]
  when "$"
    path = command(path, line.delete_prefix("$ "))
  else
    parse_contents(path, line, files)
  end
  total += 1
end

total = 0
files.each do |key, bytes|
  total += bytes if bytes < 100000
end

puts "Part 1: #{total}"

need = files["/"] - 40000000
free = files["/"]
files.each_value do |bytes|
  free = [bytes, free].min if bytes >= need
end
puts "Part 2: #{free}"
