total = 0

input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  puts "#{index}: #{line}"
  total += 1
end

puts
puts "#{total} lines"
