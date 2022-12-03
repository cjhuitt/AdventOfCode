def outcome(result)
  return 0 if result == "X"
  return 3 if result == "Y"
  return 6 if result == "Z"
end

def translate(them, result)
  return "C" if them == "A" and result == "X"
  return "A" if them == "A" and result == "Y"
  return "B" if them == "A" and result == "Z"
  return "A" if them == "B" and result == "X"
  return "B" if them == "B" and result == "Y"
  return "C" if them == "B" and result == "Z"
  return "B" if them == "C" and result == "X"
  return "C" if them == "C" and result == "Y"
  return "A" if them == "C" and result == "Z"
end

def hand(me)
  return 1 if me == "A"
  return 2 if me == "B"
  return 3 if me == "C"
end

sum = 0
File.foreach("input.txt").with_index do |line, index|
  them, result = line.split
  sum += outcome(result) + hand(translate(them, result))
end

puts "Total: #{sum}"
