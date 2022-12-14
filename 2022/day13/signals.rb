$pattern = /\[([^\[\]]*)\]/

def compare(first, second, subs)
  first_parts = first.split(",")
  second_parts = second.split(",")
  first_parts.zip(second_parts).each do |pair|
    a = pair[0]
    b = pair[1]
    return -1 if a.nil?
    return 1 if b.nil?
    comp = 0
    if subs.has_key? a or subs.has_key? b
      a = subs[a] if subs.has_key? a
      b = subs[b] if subs.has_key? b
      comp = compare(a, b, subs)
    else
      comp = a.to_i.<=> b.to_i
    end
    return comp if comp != 0
  end
  first_parts.length.<=>(second_parts.length)
end

def less?(first, second)
  substr = "a"
  subs = {}
  while first.match? $pattern
    first.gsub!($pattern) do |match|
      sub = substr
      substr.next!
      subs[sub] = $1
      sub
    end
  end
  while second.match? $pattern
    second.gsub!($pattern) do |match|
      sub = substr
      substr.next!
      subs[sub] = $1
      sub
    end
  end

  compare(first, second, subs) == -1
end

sum = 0
index = 0
input = ARGV.fetch(0, "input.txt")
lines = File.readlines(input, chomp: true)
until lines.empty?
  signals = lines.shift(3)
  index += 1
  sum += index if less?(signals[0], signals[1])
end

puts "\n#{index} pairs"
puts "Sum of correct order indices: #{sum}"
