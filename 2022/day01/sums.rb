total = 0
all_totals = Array.new

File.foreach("input.txt") do |line|
  unless line.strip.empty?
    total += line.to_i
  else
    all_totals.push total
    total = 0
  end
end

all_totals.sort!

puts "1 elf max calories: #{all_totals.last}"
puts "3 elf max calories: #{all_totals.last(3).sum}"
