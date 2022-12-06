require 'set'

def start(line, range)
  index = range
  while index < (line.size) do
    break if line[index - range,range].split("").uniq.size == range
    index += 1
  end
  index
end

line = File.readlines("input.txt", chomp: true).join()
puts "Signal starts at #{start(line, 4)}"
puts "Message starts at #{start(line, 14)}"
