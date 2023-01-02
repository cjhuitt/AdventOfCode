def to_decimal(input)
  input.reverse.chars.collect.with_index do |c, position|
    case c
    when "0"
      0
    when "1"
      1
    when "2"
      2
    when "-"
      -1
    when "="
      -2
    end * (5 ** position)
  end.sum
end

def to_snafu(val)
  removed = 0
  digits = []
  until val == 0
    r = val % 5
    val /= 5
    val += 1 if r >= 3
    digits << ["0", "1", "2", "=", "-"][r]
  end
  digits.reverse.join
end

sum = 0
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true) do |line|
  sum += to_decimal(line)
end

puts "Values sum to #{sum} (#{to_snafu(sum)})"
